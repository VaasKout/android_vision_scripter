package com.vision.scripter.data.impl

import com.vision.scripter.data.api.ControlStreamer
import com.vision.scripter.data.api.CvStreamer
import com.vision.scripter.data.api.ScripterRepository
import com.vision.scripter.data.api.VideoStreamer
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent

@Module
@InstallIn(SingletonComponent::class)
interface DataBindModule {

    @Binds
    fun bindScripterRepository(scripterRepositoryImpl: ScripterRepositoryImpl): ScripterRepository

    @Binds
    fun bindVideoStreamer(videoStreamerImpl: VideoStreamerImpl): VideoStreamer

    @Binds
    fun bindCvStreamer(cvStreamerImpl: CvStreamerImpl): CvStreamer

    @Binds
    fun bindControlStreamer(controlStreamerImpl: ControlStreamerImpl): ControlStreamer
}