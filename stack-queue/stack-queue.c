#include "stack-queue.h"
#include <string.h>
#include <stdlib.h>

#define INIT_ARR_SIZE 8

struct stack {
	size_t len, cap;
	void **v;
};

Stack *mkStack(void)
{
	Stack *s = malloc(sizeof(*s));
	s->len = 0;
	s->cap = INIT_ARR_SIZE;
	s->v = calloc(s->cap, sizeof(*(s->v)));
	return s;
}

size_t stack_len(Stack *s)
{
	return s->len;
}

void *stack_peek(Stack *s)
{
	if (s->len == 0)
		return NULL;
	return s->v[s->len - 1];
}

void *stack_pop(Stack *s)
{
	void *v, *tmp;
	if (s->len == 0)
		return NULL;
	v = s->v[s->len - 1];
	s->len--;
	if (s->len < s->cap / 4) {
		s->cap /= 2;
		realloc(s->v, s->cap * sizeof(*(s->v)));
	}
	return v;
}

void stack_push(Stack *s, void *v)
{
	if (s->len == s->cap) {
		s->cap *= 2;
		realloc(s->v, s->cap * sizeof(*(s->v)));
	}
	s->v[s->len++] = v;
}

struct queue {
	size_t cap, head, tail; /* tail points to immediately after data's end */
	void **v;
};

Queue *mkQueue(void)
{
	Queue *q = malloc(sizeof(*q));
	q->cap = INIT_ARR_SIZE;
	q->v = calloc(q->cap, sizeof(*(q->v)));
	q->tail = q->head = 0;
	return q;
}

size_t queue_len(Queue *q)
{
	return q->tail - q->head;
}

void *queue_peek(Queue *q)
{
	if (q->head == q->tail)
		return NULL;
	return q->v[q->head];
}

void *queue_pop(Queue *q)
{
	void *v;
	if (q->head == 0)
		return NULL;
	v = q->v[q->head++];
	if (q->head > INIT_ARR_SIZE && q->head > q->tail - q->head) {
		memcpy(q->v, q->v + q->head, q->tail - q->head);
	}
	if (q->tail < q->cap / 4) {
		q->cap /= 2;
		realloc(q->v, q->cap * sizeof(*(q->v)));
	}
	return v;
}
void queue_push(Queue *q, void *v)
{
	if (q->tail == q->cap) {
		q->cap *= 2;
		realloc(q->v, q->cap * sizeof(*(q->v))); 
	}
	q->v[q->tail++] = v;
}


//TODO: finish impl
