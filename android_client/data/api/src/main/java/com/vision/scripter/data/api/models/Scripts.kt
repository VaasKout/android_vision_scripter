package com.vision.scripter.data.api.models

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class SaveStepRequest(
    @SerialName("serial")
    val serial: String,
    @SerialName("name")
    val name: String,
    @SerialName("step")
    val scriptStep: ScriptStep,
)

@Serializable
data class Script(
    @SerialName("name")
    val name: String = "",
    @SerialName("steps")
    val steps: List<ScriptStep> = listOf(),
)

const val APPLY_ON_TEMPLATE = "template"
const val APPLY_ON_TEXT = "text"

@Serializable
data class ScriptStep(
    @SerialName("events")
    val events: List<StepEvent> = listOf(),
    @SerialName("action")
    val action: String = "",
    @SerialName("template")
    val template: Boolean,
    @SerialName("text")
    val text: String = "",
    @SerialName("command")
    val command: String = "",
)

fun ScriptStep.isEmpty(): Boolean {
    return events.isEmpty() && action.isEmpty() && !template && text.isEmpty() && command.isEmpty()
}

@Serializable
data class StepEvent(
    @SerialName("time")
    val time: Long = 0L,
    @SerialName("data")
    val data: ByteArray? = null,
)

@Serializable
data class KeyboardButtons(
    @SerialName("buttons")
    val buttons: List<RectangleWithText> = listOf(),
)