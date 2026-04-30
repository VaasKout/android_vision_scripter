package com.vision.scripter.welcome.impl.ui

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Scaffold
import androidx.compose.material3.SnackbarHostState
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalSoftwareKeyboardController
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import com.vision.scripter.ui.CustomButton
import com.vision.scripter.ui.ProvideSnackbarHost
import com.vision.scripter.welcome.impl.R

@Composable
fun WelcomeScreen(
    uiStateHolder: WelcomeUiStateHolder,
    snackbarHostState: SnackbarHostState,
) {
    val screenUiState = uiStateHolder.uiStateFlow.collectAsStateWithLifecycle(
        WelcomeUiState()
    ).value

    LaunchedEffect(Unit) {
        uiStateHolder.onInitData()
    }

    val keyboard = LocalSoftwareKeyboardController.current

    Scaffold(
        modifier = Modifier.fillMaxSize(),
        snackbarHost = {
            ProvideSnackbarHost(snackbarHostState = snackbarHostState)
        },
    ) {
        Box(
            modifier = Modifier
                .fillMaxSize()
                .padding(it)
        ) {
            Column(
                modifier = Modifier
                    .align(Alignment.Center)
                    .padding(horizontal = 8.dp)
            ) {
                UrlEditText(
                    modifier = Modifier.fillMaxWidth(),
                    url = screenUiState.url,
                    onValueChange = {
                        uiStateHolder.editUrl(it)
                    }
                )
                PortEditText(
                    modifier = Modifier.fillMaxWidth(),
                    port = screenUiState.port,
                    onValueChange = {
                        uiStateHolder.editPort(it)
                    }
                )
                CustomButton(
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 8.dp),
                    text = stringResource(R.string.save_button),
                    onClick = {
                        keyboard?.hide()
                        uiStateHolder.onApplyData()
                    }
                )
            }
        }
    }
}

@Composable
fun UrlEditText(
    modifier: Modifier = Modifier,
    url: String,
    onValueChange: (String) -> Unit,
) {
    OutlinedTextField(
        modifier = modifier,
        value = url,
        onValueChange = {
            onValueChange(it)
        },
        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Uri),
        label = {
            Text(
                text = stringResource(R.string.url_hint)
            )
        },
        singleLine = true,
    )
}

@Composable
fun PortEditText(
    modifier: Modifier = Modifier,
    port: String,
    onValueChange: (String) -> Unit,
) {
    OutlinedTextField(
        modifier = modifier,
        value = port,
        onValueChange = {
            onValueChange(it)
        },
        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.NumberPassword),
        label = {
            Text(
                text = stringResource(R.string.port_hint)
            )
        },
        singleLine = true,
    )
}