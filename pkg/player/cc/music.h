#ifndef HEADER_CONTROLLER_H
#define HEADER_CONTROLLER_H
void launchFMOD();
void preparePlayingFMOD();
void setMediaFMOD(char* path);
void playFMOD();
void pauseFMOD();
void stopFMOD();
void exitFMOD();
int getPlayingFMOD();
unsigned int getLengthFMOD();
unsigned int getPositionFMOD();
void setPositionFMOD(unsigned int ms);
void setVolumeFMOD(float volume);
void setFrequencyFMOD(float frequency);
float getPitchFMOD();
void setPitchFMOD(float pitch);
void setPitchDspFMOD(float pitch);
#endif