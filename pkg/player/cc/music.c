#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "music.h"
#include "inc/fmod.h"
//#include "fmod_common.h"

#define FMOD_FALSE 0
#define FMOD_TRUE 1

// FMOD_System_Create and Init args.
FMOD_SYSTEM* f_system = NULL;
void* f_system_extra_data = NULL;

// Controller
int playing = FMOD_FALSE;
int pausing = FMOD_TRUE;

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
    stopFMOD();
    FMOD_System_CreateSound(f_system, path, FMOD_DEFAULT, NULL, &f_sound);
    FMOD_System_PlaySound(f_system, f_sound, NULL, 1, &f_channel);
    FMOD_Channel_GetFrequency(f_channel, &frequency_base);
    FMOD_Channel_SetVolume(f_channel, 1);
}

void playFMOD() {
    pausing = FMOD_FALSE;
    if (!playing) {
        playing = FMOD_TRUE;
        while(playing) {
            FMOD_Channel_SetPaused(f_channel, pausing);
            FMOD_System_Update(f_system);
            usleep(100 * 1000);
            if (!playing) {
                break;
            }
            if (!pausing) {
                FMOD_Channel_IsPlaying(f_channel, &playing);
                if (!playing) {
                    FMOD_System_PlaySound(f_system, f_sound, NULL, 1, &f_channel);
                }
            }
        }
    }
}

void pauseFMOD() {
    playing = FMOD_FALSE;
    pausing = FMOD_TRUE;
    FMOD_Channel_SetPaused(f_channel, pausing);
    FMOD_System_Update(f_system);
}

void stopFMOD() {
    if (NULL != f_sound) {
        FMOD_Sound_Release(f_sound);
        if (NULL != f_channel) {
            FMOD_Channel_Stop(f_channel);
            if (NULL != f_channel) {
                printf("f_channel is not nullptr.\n");
            } // f_channel = NULL;
        }
    }
}

void exitFMOD() {
    stopFMOD();
    if (NULL != f_system) {
        FMOD_System_Release(f_system);
    }
}

int getPlayingFMOD() {
    return playing;
}

unsigned int getLengthFMOD() {
    if (NULL != f_sound) {
        unsigned int ms;
        FMOD_Sound_GetLength(f_sound, &ms, FMOD_TIMEUNIT_MS);
        return ms;
    }
    return 0;
}

unsigned int getPositionFMOD() {
    if (NULL != f_channel) {
        unsigned int ms;
        FMOD_Channel_GetPosition(f_channel, &ms, FMOD_TIMEUNIT_MS);
        return ms;
    }
    return 0;
}

void setPositionFMOD(unsigned int ms) {
    if (NULL != f_channel) {
        FMOD_Channel_SetPosition(f_channel, ms, FMOD_TIMEUNIT_MS);
    }
}
