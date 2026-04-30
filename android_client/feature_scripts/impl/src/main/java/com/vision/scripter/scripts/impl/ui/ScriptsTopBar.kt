package com.vision.scripter.scripts.impl.ui

import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import com.vision.scripter.ui.TopBar
import com.vision.scripter.ui.customClickable

@Composable
internal fun ScriptsTopBar(
    onBackClick: () -> Unit
) {
    TopBar(
        startContent = {
            Icon(
                modifier = Modifier.customClickable(
                    onClick = onBackClick,
                ),
                imageVector = Icons.AutoMirrored.Filled.ArrowBack,
                tint = MaterialTheme.colorScheme.onSurface,
                contentDescription = ""
            )
        },
        endContent = {},
    )
}

@Preview
@Composable
private fun ControlTopBarPreview() {
    ScriptsTopBar(
        onBackClick = {}
    )
}