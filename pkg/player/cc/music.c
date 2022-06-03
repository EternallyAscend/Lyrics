#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "music.h"
#include "inc/fmod.h"
//#include "fmod_common.h"

// FMOD_System_Create and Init args.
FMOD_SYSTEM* f_system = NULL;
void* f_system_extra_data = NULL;

// Controller
int playing = 0;
int pausing = 1;

// 基数
float frequency_base;

FMOD_SOUND* f_sound = NULL;
FMOD_CHANNEL* f_channel = NULL;

void launchFMOD() {
    if (NULL == f_system) {
        FMOD_System_Create(&f_system, FMOD_VERSION);
        FMOD_System_Init(f_system, 256, FMOD_INIT_NORMAL, f_system_extra_data);
    }
}

void setMediaFMOD(char* path) {
    printf("%s\n", path);
    FMOD_System_CreateSound(f_system, path, FMOD_DEFAULT, NULL, &f_sound);
    FMOD_System_PlaySound(f_system, f_sound, NULL, 1, &f_channel);
    FMOD_Channel_GetFrequency(f_channel, &frequency_base);
    FMOD_Channel_SetVolume(f_channel, 1);
}

void playFMOD() {
    pausing = 0;
    playing = 1;
    while(playing) {
        FMOD_Channel_SetPaused(f_channel, pausing);
        FMOD_System_Update(f_system);
        usleep(100 * 1000);
        if (!playing) {
            break;
        }
        if (!pausing) {
            FMOD_Channel_IsPlaying(f_channel, &playing);
        }
    }
}

void pauseFMOD() {
    pausing = 1;
    playing = 0;
}

void stopFMOD() {
    if (NULL != f_sound) {
        FMOD_Sound_Release(f_sound);
        if (NULL != f_channel) {
            f_channel = NULL;
        }
    }
}

void exitFMOD() {
    if (NULL != f_system) {
        FMOD_System_Release(f_system);
    }
}
