package com.vision.scripter.streaming.impl.ui.items.menu

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.filled.Search
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.vision.scripter.ui.customClickable

@Composable
fun KeyboardMenu(
    modifier: Modifier = Modifier,
    isLoading: Boolean,
    onKeyboardInitClick: () -> Unit,
    onBackClick: () -> Unit,
) {
    Column(
        modifier = modifier
            .background(
                color = Color.White,
                shape = RoundedCornerShape(topStart = 8.dp, bottomStart = 8.dp),
            )
            .padding(horizontal = 4.dp),
        horizontalAlignment = Alignment.End,
        verticalArrangement = Arrangement.spacedBy(4.dp),
    ) {

        if (isLoading) {
            CircularProgressIndicator(modifier = Modifier.size(32.dp))
        } else {
            Icon(
                modifier = Modifier
                    .size(32.dp)
                    .customClickable(onClick = onKeyboardInitClick),
                imageVector = Icons.Filled.Search,
                tint = MaterialTheme.colorScheme.onSurface,
                contentDescription = ""
            )
        }

        Icon(
            modifier = Modifier
                .size(32.dp)
                .customClickable(onClick = onBackClick),
            imageVector = Icons.AutoMirrored.Filled.ArrowBack,
            tint = MaterialTheme.colorScheme.onSurface,
            contentDescription = ""
        )
    }
}

@Preview
@Composable
private fun KeyboardMenuDefaultPreview() {
    KeyboardMenu(
        onKeyboardInitClick = {},
        isLoading = false,
        onBackClick = {},
    )
}
