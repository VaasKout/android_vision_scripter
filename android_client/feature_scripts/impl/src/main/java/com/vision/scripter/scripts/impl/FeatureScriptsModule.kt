package com.vision.scripter.scripts.impl

import com.vision.scripter.scripts.api.FeatureScripts
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.android.components.ActivityComponent

@Module
@InstallIn(ActivityComponent::class)
interface FeatureScriptsModule {

    @Binds
    fun bindScriptsScreen(featureScriptsImpl: FeatureScriptsImpl): FeatureScripts
}