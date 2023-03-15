#include "hal.h"

#define DELAY 500
#define DELAY_STEP 150

#define SWITCHES_HANDLE_DELAY 100

int leds[] = {GPIO_PIN_3, GPIO_PIN_4, GPIO_PIN_5,
                GPIO_PIN_6, GPIO_PIN_8, GPIO_PIN_9,
                GPIO_PIN_11, GPIO_PIN_12};
int switches[] = {GPIO_PIN_4, GPIO_PIN_8, GPIO_PIN_10, GPIO_PIN_12};

void resetLeds();
void setFrame(size_t ledIndex);
void setFrame(size_t startLedIndex, size_t lastLedIndex);

int currentFrame = sizeof(leds)/sizeof(int)-1;

int addedDelay = 0;

void TIM6_IRQ_Handler()
{
    resetLeds();
    if (currentFrame == 0) 
    {
        setFrame(0);
        setFrame(sizeof(leds)/sizeof(int)-1);
        currentFrame = sizeof(leds)/sizeof(int)-1;
    }
    else 
    {
        setFrame(currentFrame-1, currentFrame);
        currentFrame--;
    }
}

void TIM7_IRQ_Handler()
{
    __disable_irq();
    addedDelay = 0;
    for (size_t i=0; i<sizeof(switches)/sizeof(int); i++) 
    {
        addedDelay*=2;
        GPIO_PinState state = HAL_GPIO_ReadPin(GPIOE, switches[i]);
        if (state == GPIO_PIN_SET)
            addedDelay++;
    }
    WRITE_REG(TIM6_ARR, DELAY+addedDelay*DELAY_STEP);
    WRITE_REG(TIM6_DIER, TIM_DIER_UIE);
    WRITE_REG(TIM6_PSC, 0);
    
    WRITE_REG(TIM6_CR1, TIM_CR1_CEN);
    __enable_irq();
}

int umain() 
{
    registerTIM6_IRQHandler(TIM6_IRQ_Handler);
    registerTIM7_IRQHandler(TIM7_IRQ_Handler);

    WRITE_REG(TIM6_ARR, DELAY);
    WRITE_REG(TIM6_DIER, TIM_DIER_UIE);
    WRITE_REG(TIM6_PSC, 0);

    WRITE_REG(TIM7_ARR, SWITCHES_HANDLE_DELAY);
    WRITE_REG(TIM7_DIER, TIM_DIER_UIE);
    WRITE_REG(TIM7_PSC, 0);

    WRITE_REG(TIM6_CR1, TIM_CR1_CEN);
    WRITE_REG(TIM7_CR1, TIM_CR1_CEN);

    __enable_irq();

    return 0;
}

void resetLeds() 
{
    for (size_t i = 0; i < sizeof(leds)/sizeof(int); i++) 
    {
        HAL_GPIO_WritePin(GPIOD, leds[i], GPIO_PIN_RESET);
    }
}

void setFrame(size_t ledIndex) 
{
    HAL_GPIO_WritePin(GPIOD, leds[ledIndex], GPIO_PIN_SET);
}

void setFrame(size_t startLedIndex, size_t lastLedIndex) 
{
    for (size_t i=startLedIndex; i<=lastLedIndex; i++) 
        setFrame(i);
}