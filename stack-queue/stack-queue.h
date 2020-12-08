#ifndef STACK_QUEUE_H
#define STACK_QUEUE_H

#include <stddef.h>

typedef struct stack Stack;
typedef struct queue Queue;

Stack *mkStack(void);

size_t stack_len(Stack *s);
void *stack_peek(Stack *s);
void *stack_pop(Stack *s);
void stack_push(Stack *s, void *v);

Queue *mkQueue(void);

size_t queue_len(Queue *q);
void *queue_peek(Queue *q);
void *queue_pop(Queue *q);
void queue_push(Queue *q, void *v);

#endif
