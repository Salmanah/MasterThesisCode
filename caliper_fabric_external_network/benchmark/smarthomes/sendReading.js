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
    if (bc.bcType === 'fabric-ccp') {
        args = {
            chaincodeFunction: 'sendDeviceReading',
            chaincodeArguments: ['DEVICE_001','FRIDGE','67 98 fd 90 bd 92 a1 11 5f df a2 be 11 57 e5 90 23 2d cf 76 7a 0b 11 fd 60 1a 66 b2 72 6d bf 67 69 1e 42 c2 0d fe 78 76 14 01 7c b9 d3 94 70 1f 5e 0a c0 4b 12 f8 58 19 f5 a5 43 19 49 e3 9e 9e be d9 04 6b 92 a4 22 08 98 7f 32 f1 73 4a 48 ac f4 4a 8d 2a 9f 6c fe e6 de 36 13 b1 84 96 94 c1 6e 96 44 21 89 02 3a ea a8 95 1d 39 57 9d 7f e3 8c 65 6f 10 a4 ea 5b ab da f4 84 15 b2 d2 e2 cf 93 50 58 04 66 64 46 f8 c1 dc ec cd 96 ca e6 5c 54 9f 1ddf 55 13 c3 73 fe be 61 aa 01 04 fe 3a 75 00 67 29 e2 92 8c 8a 3a 65 ce 3f 95 9a b4 ac 75 2a 82 93 57 ed 3d f4 14 9c bb 8c d4 55 a0 7f 74 40 d6 ac d1 7f 7c ab 98 76 76 97 77 9f db 34 cc ce e1 b1 f1 a5 ba ab 79 81 f7 20 77 2d 89 df 55 d7 ee d8 65 de d8 69 28 18 a1 68 87 5e fa 74 ec f6 e1 9e 4f 3e 60 e2 28 15 bd 7c 13 cf eb c5 24 0c ef fa 2f 81 ea 67 2e ae f2 19 e2 25 89 d8 8e 18 4c 15 50 24 0f 30 b4 ec 99 15 1a 08 de 0e d1 78 14 6f 6e 5f a6 87 69 1c ca c0 a0 a1 6d 74 39 4d d0 ae aa bf 11 61 55 28 e0 ce ab 71 3f cf 15 4d dc 77 76 f2 4d 27 bc a9 f0 af c3 8c 82 5e 78 5b 77 26 4c a6 29 d7 32 d3 ac 36 9f 08 30 ae 7d 61 0d 9a 96 35 fa 64 64 40 c4 ab 74 37 59 ef a0 62 4a 2c 8a 3c 45 63 a7 af 89 03 7a d8 33 f3 41 f3 6b 3b d7 73 62 86 2d 23 e5 de 88 e8 44 12 cd b1 12 1f 8e 53 a9 78 7c aa 8b b3 79 98 d8 29 75 e7 3b bb 18 f3 74 63 63 e7 c1 e6 1e 31 8a 75 ee cb 61 d6 58 c9 be 67 7e 06 14 35 27 55 97 61 a1 70 39 0a fd 49 14 9d 0b 79 df 5a 9a 80 70 24 4d 36 29 6d d8 ae a0 95 e6 31 c2 1a b0 66 08 af 22 80 d9 6d a0 43 84 7c 13 68 df 45 e9 eb a2 1a 30 05 19 e6 f9 bb 18 2c 43 32 3a 61 c1 5f 5c']
        };
    }else{
        args = {
            verb: 'sendDeviceReading',
            ID: "DEVICE_001",
            DeviceType: "Fridge",
            Data: 'DATA'
        };
    }
    return bc.invokeSmartContract(contx, 'device', 'v0', args, 120);
};

module.exports.end = function() {
    return Promise.resolve();
};
