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
            chaincodeArguments: [deviceId,'FRIDGE','67 98 fd 90 bd 92 a1 11 5f df a2 be 11 57 e5 90 23 2d cf 76 7a 0b 11 fd 60 1a 66 b2 72 6d bf 67 69 1e 42 c2 0d fe 78 76 14 01 7c b9 d3 94 70 1f 5e 0a c0 4b 12 f8 58 19 f5 a5 43 19 49 e3 9e 9e be d9 04 6b 92 a4 22 08 98 7f 32 f1 73 4a 48 ac f4 4a 8d 2a 9f 6c fe e6 de 36 13 b1 84 96 94 c1 6e 96 44 21 89 02 3a ea a8 95 1d 39 57 9d 7f e3 8c 65 6f 10 a4 ea 5b ab da f4 84 15 b2 d2 e2 cf 93 50 58 04 66 64 46 f8 c1 dc ec cd 96 ca e6 5c 54 9f 1ddf 55 13 c3 73 fe be 61 aa 01 04 fe 3a 75 00 67 29 e2 92 8c 8a 3a 65 ce 3f 95 9a b4 ac 75 2a 82 93 57 ed 3d f4 14 9c bb 8c d4 55 a0 7f 74 40 d6 ac d1 7f 7c ab 98 76 76 97 77 9f db 34 cc ce e1 b1 f1 a5 ba ab 79 81 f7 20 77 2d 89 df 55 d7 ee d8 65 de d8 69 28 18 a1 68 87 5e fa 74 ec f6 e1 9e 4f 3e 60 e2 28 15 bd 7c 13 cf eb c5 24 0c ef fa 2f 81 ea 67 2e ae f2 19 e2 25 89 d8 8e 18 4c 15 50 24 0f 30 b4 ec 99 15 1a 08 de 0e d1 78 14 6f 6e 5f a6 87 69 1c ca c0 a0 a1 6d 74 39 4d d0 ae aa bf 11 61 55 28 e0 ce ab 71 3f cf 15 4d dc 77 76 f2 4d 27 bc a9 f0 af c3 8c 82 5e 78 5b 77 26 4c a6 29 d7 32 d3 ac 36 9f 08 30 ae 7d 61 0d 9a 96 35 fa 64 64 40 c4 ab 74 37 59 ef a0 62 4a 2c 8a 3c 45 63 a7 af 89 03 7a d8 33 f3 41 f3 6b 3b d7 73 62 86 2d 23 e5 de 88 e8 44 12 cd b1 12 1f 8e 53 a9 78 7c aa 8b b3 79 98 d8 29 75 e7 3b bb 18 f3 74 63 63 e7 c1 e6 1e 31 8a 75 ee cb 61 d6 58 c9 be 67 7e 06 14 35 27 55 97 61 a1 70 39 0a fd 49 14 9d 0b 79 df 5a 9a 80 70 24 4d 36 29 6d d8 ae a0 95 e6 31 c2 1a b0 66 08 af 22 80 d9 6d a0 43 84 7c 13 68 df 45 e9 eb a2 1a 30 05 19 e6 f9 bb 18 2c 43 32 3a 61 c1 5f 5c']
        };
    }else{
        args = {
            verb: 'sendDeviceReading',
            id: deviceId,
            deviceType: "Thermo",
            data: 'd8 d2 fe 06 2e 84 4e 8e 6c a5 4c b8 96 8e c4 9d 91 99 ea 8e 40 20 9d f6 eb 35 c7 c1 9c 13 09 6a c2 25 5a 92 8b 72 ab fe b8 3d 6f 6d e8 a5 11 9e ae a4 32 41 eb 60 aa 5b 31 6b 41 73 ba d4 46 f4 a9 46 06 86 6c 4c 25 dc 5e 0e 08 34 b8 98 0e 99 a5 bd 6f 88 8e b2 20 33 f5 53 ad 13 be ae 81 5f be 52 69 35 73 2e 45 7d f4 6c 61 8c a8 1d c5 3d 1a 2d 4b ba e8 ce ca 38 79 62 1a 57 47 91 8d ac 99 c2 31 42 c5 ae c0 ea a8 6e d0 8e 9e 32 e9 02 58 06 54 82 98 e8 1d 4d 18 88 7e 68 79 70 bd c5 3e b4 6f 32 d7 c4 9e c2 65 6f 86 93 b3 ce 52 36 e4 7e e8 73 c2 6a b4 50 14 f6 50 d8 43 61 1f 4e e5 23 92 66 62 94 77 12 c4 c8 a6 50 b4 52 44 43 62 0a 58 0f 53 34 10 8a 92 f9 a8 ab 48 0c de 83 24 6b be ff e0 9e 66 96 6e f4 9b f1 32 71 b7 8b c4 dd c9 a6 f6 83 bb 2e 86 01 d5 ba 2e 1b 72 a2 59 94 33 26 c7 df f9 3b 08 e2 73 24 2d 2a 91 13 d4 93 b0 7f 0a 4f 4c a4 57 5f 7d 7a e1 15 88 54 f7 25 b0 35 65 d6 70 e4 b5 59 38 df 79 11 58 f3 81 b4 11 d4 ae 84 76 73 35 76 1d 88 3a 56 78 3d 18 fa 61 a6 61 33 32 48 04 0c 08 0f 93h h'
        };
    }
    return bc.invokeSmartContract(contx, 'device', 'v0', args, 120);
};

module.exports.end = function() {
    return Promise.resolve();
};

/* 1013 charcters
d8 d2 fe 06 2e 84 4e 8e 6c a5 4c b8 96 8e c4 9d 91 99 ea 8e 40 20 9d f6 eb 35 c7 c1 9c 13 09 6a c2 25 5a 92 8b 72 ab fe b8 3d 6f 6d e8 a5 11 9e ae a4 32 41 eb 60 aa 5b 31 6b 41 73 ba d4 46 f4 a9 46 06 86 6c 4c 25 dc 5e 0e 08 34 b8 98 0e 99 a5 bd 6f 88 8e b2 20 33 f5 53 ad 13 be ae 81 5f be 52 69 35 73 2e 45 7d f4 6c 61 8c a8 1d c5 3d 1a 2d 4b ba e8 ce ca 38 79 62 1a 57 47 91 8d ac 99 c2 31 42 c5 ae c0 ea a8 6e d0 8e 9e 32 e9 02 58 06 54 82 98 e8 1d 4d 18 88 7e 68 79 70 bd c5 3e b4 6f 32 d7 c4 9e c2 65 6f 86 93 b3 ce 52 36 e4 7e e8 73 c2 6a b4 50 14 f6 50 d8 43 61 1f 4e e5 23 92 66 62 94 77 12 c4 c8 a6 50 b4 52 44 43 62 0a 58 0f 53 34 10 8a 92 f9 a8 ab 48 0c de 83 24 6b be ff e0 9e 66 96 6e f4 9b f1 32 71 b7 8b c4 dd c9 a6 f6 83 bb 2e 86 01 d5 ba 2e 1b 72 a2 59 94 33 26 c7 df f9 3b 08 e2 73 24 2d 2a 91 13 d4 93 b0 7f 0a 4f 4c a4 57 5f 7d 7a e1 15 88 54 f7 25 b0 35 65 d6 70 e4 b5 59 38 df 79 11 58 f3 81 b4 11 d4 ae 84 76 73 35 76 1d 88 3a 56 78 3d 18 fa 61 a6 61 33 32 48 04 0c 08 0f 93h h
*/


/* 1513 characters
67 98 fd 90 bd 92 a1 11 5f df a2 be 11 57 e5 90 23 2d cf 76 7a 0b 11 fd 60 1a 66 b2 72 6d bf 67 69 1e 42 c2 0d fe 78 76 14 01 7c b9 d3 94 70 1f 5e 0a c0 4b 12 f8 58 19 f5 a5 43 19 49 e3 9e 9e be d9 04 6b 92 a4 22 08 98 7f 32 f1 73 4a 48 ac f4 4a 8d 2a 9f 6c fe e6 de 36 13 b1 84 96 94 c1 6e 96 44 21 89 02 3a ea a8 95 1d 39 57 9d 7f e3 8c 65 6f 10 a4 ea 5b ab da f4 84 15 b2 d2 e2 cf 93 50 58 04 66 64 46 f8 c1 dc ec cd 96 ca e6 5c 54 9f 1ddf 55 13 c3 73 fe be 61 aa 01 04 fe 3a 75 00 67 29 e2 92 8c 8a 3a 65 ce 3f 95 9a b4 ac 75 2a 82 93 57 ed 3d f4 14 9c bb 8c d4 55 a0 7f 74 40 d6 ac d1 7f 7c ab 98 76 76 97 77 9f db 34 cc ce e1 b1 f1 a5 ba ab 79 81 f7 20 77 2d 89 df 55 d7 ee d8 65 de d8 69 28 18 a1 68 87 5e fa 74 ec f6 e1 9e 4f 3e 60 e2 28 15 bd 7c 13 cf eb c5 24 0c ef fa 2f 81 ea 67 2e ae f2 19 e2 25 89 d8 8e 18 4c 15 50 24 0f 30 b4 ec 99 15 1a 08 de 0e d1 78 14 6f 6e 5f a6 87 69 1c ca c0 a0 a1 6d 74 39 4d d0 ae aa bf 11 61 55 28 e0 ce ab 71 3f cf 15 4d dc 77 76 f2 4d 27 bc a9 f0 af c3 8c 82 5e 78 5b 77 26 4c a6 29 d7 32 d3 ac 36 9f 08 30 ae 7d 61 0d 9a 96 35 fa 64 64 40 c4 ab 74 37 59 ef a0 62 4a 2c 8a 3c 45 63 a7 af 89 03 7a d8 33 f3 41 f3 6b 3b d7 73 62 86 2d 23 e5 de 88 e8 44 12 cd b1 12 1f 8e 53 a9 78 7c aa 8b b3 79 98 d8 29 75 e7 3b bb 18 f3 74 63 63 e7 c1 e6 1e 31 8a 75 ee cb 61 d6 58 c9 be 67 7e 06 14 35 27 55 97 61 a1 70 39 0a fd 49 14 9d 0b 79 df 5a 9a 80 70 24 4d 36 29 6d d8 ae a0 95 e6 31 c2 1a b0 66 08 af 22 80 d9 6d a0 43 84 7c 13 68 df 45 e9 eb a2 1a 30 05 19 e6 f9 bb 18 2c 43 32 3a 61 c1 5f 5c
*/


/*
3038 characters
d8 d2 fe 06 2e 84 4e 8e 6c a5 4c b8 96 8e c4 9d 91 99 ea 8e 40 20 9d f6 eb 35 c7 c1 9c 13 09 6a c2 25 5a 92 8b 72 ab fe b8 3d 6f 6d e8 a5 11 9e ae a4 32 41 eb 60 aa 5b 31 6b 41 73 ba d4 46 f4 a9 46 06 86 6c 4c 25 dc 5e 0e 08 34 b8 98 0e 99 a5 bd 6f 88 8e b2 20 33 f5 53 ad 13 be ae 81 5f be 52 69 35 73 2e 45 7d f4 6c 61 8c a8 1d c5 3d 1a 2d 4b ba e8 ce ca 38 79 62 1a 57 47 91 8d ac 99 c2 31 42 c5 ae c0 ea a8 6e d0 8e 9e 32 e9 02 58 06 54 82 98 e8 1d 4d 18 88 7e 68 79 70 bd c5 3e b4 6f 32 d7 c4 9e c2 65 6f 86 93 b3 ce 52 36 e4 7e e8 73 c2 6a b4 50 14 f6 50 d8 43 61 1f 4e e5 23 92 66 62 94 77 12 c4 c8 a6 50 b4 52 44 43 62 0a 58 0f 53 34 10 8a 92 f9 a8 ab 48 0c de 83 24 6b be ff e0 9e 66 96 6e f4 9b f1 32 71 b7 8b c4 dd c9 a6 f6 83 bb 2e 86 01 d5 ba 2e 1b 72 a2 59 94 33 26 c7 df f9 3b 08 e2 73 24 2d 2a 91 13 d4 93 b0 7f 0a 4f 4c a4 57 5f 7d 7a e1 15 88 54 f7 25 b0 35 65 d6 70 e4 b5 59 38 df 79 11 58 f3 81 b4 11 d4 ae 84 76 73 35 76 1d 88 3a 56 78 3d 18 fa 61 a6 61 33 32 48 04 0c 08 0f 93 2f 4d 46 a1 22 a8 ad 48 7f 68 f2 57 aa 24 c0 d3 c3 02 ad a4 37 a7 66 e1 f4 a1 e0 55 33 42 ad 5a cf c6 df 3a e1 06 b5 c1 be b2 8f d6 fc 35 cf 5d d7 36 b8 99 b0 2c 03 6f b1 6c fa b6 9a 59 a9 69 49 6b 6c f4 a6 01 cc 8b 42 37 9c 4a 34 b5 7d d8 74 3a 23 9f d6 48 90 4c 3b 39 5d 9e 77 fe 44 06 e3 0d 96 a8 9b 52 dc 35 b6 91 c3 c4 5f ac 23 35 de 48 07 22 78 e1 36 4b 70 b0 4a 0c 5a f5 3d 92 31 33 ba f5 85 ba 21 95 ea 47 79 d1 06 e3 9b 10 88 f7 d6 1a 72 a8 7d b9 0b b0 08 61 fc 92 ed 8f 1a ad be da bb 1b 4c c5 f3 a4 2a 97 95 ab de 07 f9 3e 74 9a 82 c1 8a b0 fe 63 1f cd ab 0f e1 f6 4c 91 f5 ab 6c 98 46 fb 8a ab d2 b5 3e 51 8a a2 2a a8 c1 ca 3a f3 39 54 e8 93 15 fd dc df 10 6c 51 3f 51 92 35 dd 11 13 f0 96 a4 0c 6a 62 78 23 4c 77 66 e2 33 0c 9e f9 3b 3b 1c 0b 61 18 61 5b a8 8d 3b b7 a1 05 b3 1c 9b 61 fb a4 bb ec d8 45 33 16 44 6d c2 8f fc 54 b7 4b 15 f1 9c de 1d 87 31 dd d7 3f 1c e2 d3 33 b9 fa 0a 23 e0 ac de 07 46 da ee 23 91 60 3e a3 a8 9b bf 98 a8 86 8f 92 ab 1e 6f 54 fc c6 c7 1e 60 9e ee ae 41 1c dc 29 49 a8 65 75 42 8a 9e 27 4e a6 81 35 a5 cb 53 83 a7 b3 e2 55 4e 1e ce b9 7c b0 e2 8c c2 30 8e 81 cf 41 11 59 10 8b 2f ba e2 c0 01 32 1d 30 eb 4b c4 54 65 da ee 92 ed 33 39 0b 08 2c 7d e9 51 92 23 27 06 e3 21 68 ba 94 af 78 60 55 25 1d 32 48 1f ff 60 15 fc e4 ee ce 55 95 f7 bb 43 da fb b7 9c 1d 26 76 bf 92 b8 b1 d2 b6 d1 de df 95 de 84 bb 8e 9c 85 4a e8 ec 00 f8 29 0c 41 3b 5b 94 15 cc 1b 5a 8b 2a 21 98 ed 65 9e f0 e1 bc b8 c3 fe 2a 40 9a 56 1a 0a 81 39 73 73 60 1c c2 d0 78 6f 83 e5 02 c6 8f 92 da 35 ae 30 53 a2 39 11 5b c0 04 cf 55 c9 ec d9 e5 75 03 67 07 ab dd 42 80 16 5c f4 d1 53 a8 3c 13 b6 8f 74 8f 9b 08 84 6d 6c a8 bd 58 85 0e 1a 5f ee fa 3d 5d 87 8b 01 bd e7 75 a2 73 51 3d 01 16 ba ed b7 0c ef a3 2f 80 03 85 7c 23 4e 0a db f4 5f 91 89 38 be 71 d1 c3 4c d7 35 9d d3 a3 a8 31 bc 11 66 4b 59 2a 41 4b 09 cb e1 8d fb 79 b0 3d 13 28 3b 2c 6d df 7f 27 ed a4 ff 3f 85 15 c9 4a 80 aa 7c 94 39 50 60 63 bb 21 99 2f 87 41 db 04 39 0a 9b a1 cf ce 5d 80 f1 93 63 da 6b eb fa 64 61 df 23 0f 3c 36 54 bb e7 5d df 0e 23 ef 5e bb 16
*/
