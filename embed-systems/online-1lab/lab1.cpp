#include "hal.h"

#define DELAY 500
#define CLICK_DELAY 100
#define MAGIC_MODE 0x9

int leds[] = {GPIO_PIN_3, GPIO_PIN_4, GPIO_PIN_5,
                GPIO_PIN_6, GPIO_PIN_8, GPIO_PIN_9,
                GPIO_PIN_11, GPIO_PIN_12};
int switches[] = {GPIO_PIN_4, GPIO_PIN_8, GPIO_PIN_10, GPIO_PIN_12};
bool isAnimationStopped = false;

void setFrame(size_t startLedIndex, size_t lastLedIndex);
void resetLeds();
void startAnimation();
void checkStopping();
bool isStateChanged();
void activateStopMode();
void activateAnimationMode();
void activateSwitchMode();


int umain() 
{
    activateSwitchMode();
    int switchesValue;
    while (true) 
    {
        switchesValue = 0;

        for (size_t i=0; i<sizeof(switches)/sizeof(int); i++) 
        {
            switchesValue*=2;
            GPIO_PinState state = HAL_GPIO_ReadPin(GPIOE, switches[i]);
            HAL_GPIO_WritePin(GPIOD, leds[i], state);

            if (state == GPIO_PIN_SET)
                switchesValue++;
        }

        if (switchesValue == MAGIC_MODE) 
            startAnimation();

    }

    return 0;
    
}

void startAnimation()
{
    activateAnimationMode();
    while (true) {
        checkStopping();

        for (size_t i=0; i < sizeof(leds)/sizeof(int)-2; i++) 
        {
            resetLeds();
            setFrame(i, i+2);
            checkStopping();
    
        }
        for (size_t i=sizeof(leds)/sizeof(int)-4; i > 0; i--) 
        {
            resetLeds();
            setFrame(i, i+2);
            checkStopping();
        }
    }
}

void checkStopping() 
{
    if(isStateChanged()) {
        isAnimationStopped = !isAnimationStopped;
    }

    if (isAnimationStopped) 
        activateStopMode();
    while (isAnimationStopped) 
    {
        if(isStateChanged())
            isAnimationStopped = !isAnimationStopped;
        if (!isAnimationStopped)
            activateAnimationMode();
    }
}

bool isStateChanged() 
{
    bool isStateChanged = HAL_GPIO_ReadPin(GPIOC, GPIO_PIN_15) == GPIO_PIN_RESET;

    HAL_Delay(DELAY);

    return isStateChanged;
}

void activateSwitchMode()
{
    HAL_GPIO_WritePin(GPIOD, GPIO_PIN_13, GPIO_PIN_RESET);
    HAL_GPIO_WritePin(GPIOD, GPIO_PIN_15, GPIO_PIN_RESET);
    HAL_GPIO_WritePin(GPIOD, GPIO_PIN_14, GPIO_PIN_SET);
}

void activateStopMode() 
{
    HAL_GPIO_WritePin(GPIOD, GPIO_PIN_13, GPIO_PIN_RESET);
    HAL_GPIO_WritePin(GPIOD, GPIO_PIN_14, GPIO_PIN_RESET);
    HAL_GPIO_WritePin(GPIOD, GPIO_PIN_15, GPIO_PIN_SET);
}

void activateAnimationMode()
{
    HAL_GPIO_WritePin(GPIOD, GPIO_PIN_15, GPIO_PIN_RESET);
    HAL_GPIO_WritePin(GPIOD, GPIO_PIN_14, GPIO_PIN_RESET);
    HAL_GPIO_WritePin(GPIOD, GPIO_PIN_13, GPIO_PIN_SET);
}

void setFrame(size_t startLedIndex, size_t lastLedIndex) 
{
    for (size_t i=startLedIndex; i<=lastLedIndex; i++) 
    {
        HAL_GPIO_WritePin(GPIOD, leds[i], GPIO_PIN_SET);
    }
}

void resetLeds() 
{
    for (size_t i = 0; i < sizeof(leds)/sizeof(int); i++) 
    {
        HAL_GPIO_WritePin(GPIOD, leds[i], GPIO_PIN_RESET);
    }
}

