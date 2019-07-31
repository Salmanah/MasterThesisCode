/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

'use strict';

module.exports.info  = 'Sending device reading.';

let txindex = 0;
let bc, contx;


module.exports.init = function(blockchain, context, args) {
    bc = blockchain;
    contx = context;
    return Promise.resolve();
};

module.exports.run = function() {
    let args;
    txindex++;
    let deviceId = "DEVICE_00"+txindex.toString();

    if (bc.bcType === 'fabric-ccp') {
        args = {
            chaincodeFunction: 'sendDeviceReading',
            chaincodeArguments: [deviceId,'FRIDGE','d8 d2 fe 06 2e 84 4e 8e 6c a5 4c b8 96 8e c4 9d 91 99 ea 8e 40 20 9d f6 eb 35 c7 c1 9c 13 09 6a c2 25 5a 92 8b 72 ab fe b8 3d 6f 6d e8 a5 11 9e ae a4 32 41 eb 60 aa 5b 31 6b 41 73 ba d4 46 f4 a9 46 06 86 6c 4c 25 dc 5e 0e 08 34 b8 98 0e 99 a5 bd 6f 88 8e b2 20 33 f5 53 ad 13 be ae 81 5f be 52 69 35 73 2e 45 7d f4 6c 61 8c a8 1d c5 3d 1a 2d 4b ba e8 ce ca 38 79 62 1a 57 47 91 8d ac 99 c2 31 42 c5 ae c0 ea a8 6e d0 8e 9e 32 e9 02 58 06 54 82 98 e8 1d 4d 18 dd ss ss aa ss aa ss aa ss ss aa ee ff aa']
        };
    }else{
        args = {
            verb: 'sendDeviceReading',
            id: deviceId,
            deviceType: "Thermo",
            data: 'd8 d2 fe 06 2e 84 4e 8e 6c a5 4c b8 96 8e c4 9d 91 99 ea 8e 40 20 9d f6 eb 35 c7 c1 9c 13 09 6a c2 25 5a 92 8b 72 ab fe b8 3d 6f 6d e8 a5 11 9e ae a4 32 41 eb 60 aa 5b 31 6b 41 73 ba d4 46 f4 a9 46 06 86 6c 4c 25 dc 5e 0e 08 34 b8 98 0e 99 a5 bd 6f 88 8e b2 20 33 f5 53 ad 13 be ae 81 5f be 52 69 35 73 2e 45 7d f4 6c 61 8c a8 1d c5 3d 1a 2d 4b ba e8 ce ca 38 79 62 1a 57 47 91 8d ac 99 c2 31 42 c5 ae c0 ea a8 6e d0 8e 9e 32 e9 02 58 06 54 82 98 e8 1d 4d 18 dd ss ss aa ss aa ss aa ss ss aa ee ff aa'
        };
    }
    return bc.invokeSmartContract(contx, 'device', 'v0', args, 120);
};

module.exports.end = function() {
    return Promise.resolve();
};

/* 500 charcters
d8 d2 fe 06 2e 84 4e 8e 6c a5 4c b8 96 8e c4 9d 91 99 ea 8e 40 20 9d f6 eb 35 c7 c1 9c 13 09 6a c2 25 5a 92 8b 72 ab fe b8 3d 6f 6d e8 a5 11 9e ae a4 32 41 eb 60 aa 5b 31 6b 41 73 ba d4 46 f4 a9 46 06 86 6c 4c 25 dc 5e 0e 08 34 b8 98 0e 99 a5 bd 6f 88 8e b2 20 33 f5 53 ad 13 be ae 81 5f be 52 69 35 73 2e 45 7d f4 6c 61 8c a8 1d c5 3d 1a 2d 4b ba e8 ce ca 38 79 62 1a 57 47 91 8d ac 99 c2 31 42 c5 ae c0 ea a8 6e d0 8e 9e 32 e9 02 58 06 54 82 98 e8 1d 4d 18 dd ss ss aa ss aa ss aa ss ss aa ee ff aa
*/


/* 250 characters
d8 d2 fe 06 2e 84 4e 8e 6c a5 4c b8 96 8e c4 9d 91 99 ea 8e 40 20 9d f6 eb 35 c7 c1 9c 13 09 6a c2 25 5a 92 8b 72 ab fe b8 3d 6f 6d e8 a5 11 9e ae a4 32 41 eb 60 aa 5b 31 6b 41 73 ba d4 46 f4 a9 46 06 86 6c 4c 25 dc 5e 0e 08 34 b8 98 0e 99 a5 bd 6f 8
*/


/*
100 characters
d8 d2 fe 06 2e 84 4e 8e 6c a5 4c b8 96 8e c4 9d 91 99 ea 8e 40 20 9d f6 eb 35 c7 c1 9c 13 09 6a c2
*/
