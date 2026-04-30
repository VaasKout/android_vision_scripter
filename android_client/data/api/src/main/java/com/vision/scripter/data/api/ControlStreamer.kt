package com.vision.scripter.data.api

import android.view.MotionEvent
import com.vision.scripter.data.api.models.ScreenSizes

interface ControlStreamer {
    fun initConnection(
        host: String,
        port: Int,
    ): Boolean

    fun sendControlData(
        screenSizes: ScreenSizes,
        event: MotionEvent?,
    ): ByteArray?

    fun close()
}