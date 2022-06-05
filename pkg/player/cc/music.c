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

FMOD_DSP* f_dsp_pitch = NULL;

void launchFMOD() {
    if (NULL == f_system) {
        FMOD_System_Create(&f_system, FMOD_VERSION);
        FMOD_System_Init(f_system, 256, FMOD_INIT_NORMAL, f_system_extra_data);
        FMOD_System_CreateDSPByType(f_system, FMOD_DSP_TYPE_PITCHSHIFT, &f_dsp_pitch);
    }
}

void preparePlayingFMOD() {
    FMOD_System_PlaySound(f_system, f_sound, NULL, 1, &f_channel);
    FMOD_Channel_GetFrequency(f_channel, &frequency_base);
    FMOD_Channel_SetVolume(f_channel, 1);
}

void setMediaFMOD(char* path) {
    stopFMOD();
    FMOD_System_CreateSound(f_system, path, FMOD_DEFAULT, NULL, &f_sound);
    preparePlayingFMOD();
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
                    preparePlayingFMOD();
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
//                printf("f_channel is not nullptr.\n");
                f_channel = NULL;
            }
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

void setVolumeFMOD(float volume) {
    if (NULL != f_channel) {
        FMOD_Channel_SetVolume(f_channel, volume);
    }
}

void setFrequencyFMOD(float frequency) {
    FMOD_Channel_SetFrequency(f_channel, frequency_base * frequency);
}

float getPitchFMOD() {
    if (NULL != f_channel) {
        float pitch;
        FMOD_Channel_GetPitch(f_channel, &pitch);
        return pitch;
    }
    return 0;
}

void setPitchFMOD(float pitch) {
    if (NULL != f_channel) {
        FMOD_Channel_SetPitch(f_channel, pitch);
    }
}

void setPitchDspFMOD(float pitch) {
    if (NULL != f_channel) {
        if (NULL != f_dsp_pitch) {
            FMOD_Channel_RemoveDSP(f_channel, f_dsp_pitch);
        }
        FMOD_DSP_SetParameterFloat(f_dsp_pitch, FMOD_DSP_PITCHSHIFT_PITCH, pitch);
        FMOD_Channel_AddDSP(f_channel, 0, f_dsp_pitch);
    }
}
