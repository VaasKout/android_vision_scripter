package com.vision.scripter.ui

import androidx.compose.foundation.clickable
import androidx.compose.foundation.interaction.MutableInteractionSource
import androidx.compose.material3.ripple
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier

@Composable
fun Modifier.conditional(
    flag: Boolean,
    apply: Modifier.() -> Modifier,
): Modifier {
    return if (flag) this.apply() else this
}


@Composable
fun Modifier.customClickable(
    onClick: () -> Unit
): Modifier {
    return this.clickable(
        indication = ripple(bounded = false),
        interactionSource = remember { MutableInteractionSource() },
        onClick = onClick,
    )
}
