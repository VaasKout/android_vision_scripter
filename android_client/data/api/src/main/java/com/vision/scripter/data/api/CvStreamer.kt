package com.vision.scripter.data.api

interface CvStreamer {
    suspend fun connect(
        host: String,
        port: Int,
    ): Boolean

    suspend fun readRectangles(): ByteArray?
    suspend fun sendCvMode(cvMode: Int): Boolean
    fun close()
}