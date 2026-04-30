package com.vision.scripter.prefs

import com.vision.scripter.prefs.api.DataStoreRepository
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent

@Module
@InstallIn(SingletonComponent::class)
interface DataStoreBindModule {
    @Binds
    fun bindDataStore(dataStoreRepositoryImpl: DataStoreRepositoryImpl): DataStoreRepository
}