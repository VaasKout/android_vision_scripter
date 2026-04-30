package com.vision.scripter.streaming.impl

import com.vision.scripter.streaming.api.FeatureStreaming
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.android.components.ActivityComponent


@Module
@InstallIn(ActivityComponent::class)
interface FeatureStreamingModule {

    @Binds
    fun bindStreamingScreen(featureStreamingImpl: FeatureStreamingImpl): FeatureStreaming
}