//	Copyright 2024 Mike Nguyen <hey(at)mike.ee>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// JS shim to launch the relevant binary.

const os = require('os');
const { spawnSync } = require('child_process');
const process = require('process')

function getPlatformAndArch() {
  let platform = os.platform();
  let arch = os.arch();

  // Rename win32 to windows
  if (platform === 'win32') {
    platform = 'windows';
  }

  // Rename x64 to amd64
  if (arch === 'x64') {
    arch = 'amd64';
  }

  return { platform, arch };
}

function main() {
  const { platform, arch } = getPlatformAndArch();

  let fullPath = `${__dirname}/main-${platform}-${arch}`;

  let res = spawnSync(fullPath, { stdio: 'inherit' })
  let code = res.status
  if (typeof code === 'number') {
    process.exit(code)
  }
  process.exit(1)
}

if (require.main === module) {
  main()
}
