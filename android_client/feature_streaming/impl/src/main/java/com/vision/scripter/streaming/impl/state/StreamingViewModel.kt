package com.vision.scripter.streaming.impl.state

import androidx.lifecycle.ViewModel
import com.vision.scripter.streaming.impl.ui.StreamingUiStateHolder
import dagger.hilt.android.lifecycle.HiltViewModel
import javax.inject.Inject

@HiltViewModel
class StreamingViewModel @Inject constructor(
    private val streamingInteractor: StreamingInteractor,
) : ViewModel(), StreamingUiStateHolder by streamingInteractor {

    override fun onCleared() {
        super.onCleared()
        streamingInteractor.closeStreams()
    }
}