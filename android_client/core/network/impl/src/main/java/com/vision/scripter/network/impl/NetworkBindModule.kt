package com.vision.scripter.network.impl

import com.vision.scripter.network.api.NetworkClient
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent

@Module
@InstallIn(SingletonComponent::class)
interface NetworkModule {
    @Binds
    fun bindNetworkClient(networkClientImpl: NetworkClientImpl): NetworkClient
}