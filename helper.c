typedef void (*callbackFunc) (int number);

void call_callbackFunc(callbackFunc f, int number) {
    f(number);
}