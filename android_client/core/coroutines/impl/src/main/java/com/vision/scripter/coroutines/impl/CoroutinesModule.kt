package com.vision.scripter.coroutines.impl

import com.vision.scripter.coroutines.api.CoroutineScopeFactory
import com.vision.scripter.coroutines.api.DispatchersFactory
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent

@Module
@InstallIn(SingletonComponent::class)
interface CoroutinesBindModule {

    @Binds
    fun bindDispatchersFactory(
        dispatchersFactoryImpl: DispatchersFactoryImpl
    ): DispatchersFactory

    @Binds
    fun bindCoroutineScopeFactory(
        coroutineScopeFactoryImpl: CoroutineScopeFactoryImpl,
    ): CoroutineScopeFactory
}