"""Offline regression tests for the deployment installer."""
import os
from pathlib import Path
import subprocess
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[2]


class BootstrapTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.directory = Path(self.temp.name)
        self.bin = self.directory / 'bin'
        self.bin.mkdir()
        curl = self.bin / 'curl'
        curl.write_text('''#!/bin/sh
set -eu
url=$2
file=${url##*/}
[ "$file" != "${FAIL_DOWNLOAD:-}" ] || exit 22
cp "$FIXTURES/$file" "$4"
''')
        curl.chmod(0o755)
        self.target = self.directory / 'deploy'
        self.target.mkdir()
        self.env = dict(os.environ, PATH=f'{self.bin}:{os.environ["PATH"]}',
                        FIXTURES=str(ROOT / 'example'))

    def run_installer(self, *args):
        return subprocess.run(['sh', str(ROOT / 'scripts/bootstrap-deploy.sh'), *args],
                              cwd=self.target, env=self.env, capture_output=True, text=True)

    def test_fresh_install(self):
        result = self.run_installer()
        self.assertEqual(result.returncode, 0, result.stderr)
        for file in ['compose.yaml', 'Caddyfile', '.env.example', 'README.md']:
            self.assertEqual((self.target / file).read_bytes(), (ROOT / 'example' / file).read_bytes())
        self.assertEqual((self.target / '.env').stat().st_mode & 0o777, 0o600)

    def test_existing_environment_is_preserved(self):
        (self.target / '.env').write_text('TOKEN=keep-this-secret\n')
        self.assertEqual(self.run_installer().returncode, 0)
        self.assertEqual((self.target / '.env').read_text(), 'TOKEN=keep-this-secret\n')

    def test_update_requires_force(self):
        self.assertEqual(self.run_installer().returncode, 0)
        (self.target / 'compose.yaml').write_text('custom deployment\n')
        (self.target / '.env').write_text('TOKEN=keep-this-secret\n')
        self.assertNotEqual(self.run_installer().returncode, 0)
        self.assertEqual((self.target / 'compose.yaml').read_text(), 'custom deployment\n')
        self.assertEqual(self.run_installer('--force').returncode, 0)
        self.assertEqual((self.target / '.env').read_text(), 'TOKEN=keep-this-secret\n')

    def test_failed_download_changes_nothing(self):
        (self.target / 'compose.yaml').write_text('custom deployment\n')
        self.env['FAIL_DOWNLOAD'] = 'Caddyfile'
        self.assertNotEqual(self.run_installer('--force').returncode, 0)
        self.assertEqual((self.target / 'compose.yaml').read_text(), 'custom deployment\n')
        self.assertFalse((self.target / '.env').exists())

    def test_invalid_argument_changes_nothing(self):
        self.assertEqual(self.run_installer('--unknown').returncode, 2)
        self.assertEqual(list(self.target.iterdir()), [])


if __name__ == '__main__':
    unittest.main()
