CC=gcc
SRCS=$(wildcard *.c */*.c)
OBJS=$(patsubst %.c, %.o, $(SRCS))
FLAG=-g
LIB=-lavutil -lavformat -lavcodec -lavutil -lswscale -lswresample -lSDL2
NAME=$(wildcard *.c)
TARGET=ffplayer

$(TARGET):$(OBJS)
	$(CC) -o $@ $^ $(FLAG) $(LIB)

%.o:%.c
	$(CC) -o $@ -c $< -g

clean:
	rm -rf $(TARGET) $(OBJS)
