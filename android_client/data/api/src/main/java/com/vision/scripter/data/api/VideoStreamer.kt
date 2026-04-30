package com.vision.scripter.data.api

interface VideoStreamer {
    fun connect(host: String, port: Int): Boolean
    fun readVideoHeader(): Pair<Int, Int>
    fun readFrameMeta(): Pair<Long, Int>
    fun readFrame(size: Int): ByteArray?

    fun close()
}