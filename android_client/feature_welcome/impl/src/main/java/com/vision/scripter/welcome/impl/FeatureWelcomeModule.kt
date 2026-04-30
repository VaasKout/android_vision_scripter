package com.vision.scripter.welcome.impl

import com.vision.scripter.welcome.api.FeatureWelcome
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.android.components.ActivityComponent

@Module
@InstallIn(ActivityComponent::class)
interface FeatureWelcomeModule {

    @Binds
    fun bindWelcomeScreen(featureWelcomeImpl: FeatureWelcomeImpl): FeatureWelcome
}