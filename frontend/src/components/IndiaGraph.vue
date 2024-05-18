<template>
    <main>
        <div class="file-input">
            <v-file-input label="Upload CSV" prepend-icon="mdi-upload" accept=".csv" v-model="data.selectedFile"
                ref="fileInput" chips @change="importCSV"></v-file-input>
        </div>
    </main>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { LogDebug } from '../../wailsjs/runtime'

const data = reactive({
    selectedFile: null as any,
    csv: "" as String | ArrayBuffer | null
})
const fileInput = ref<HTMLInputElement | null>(null)
const files = ref()

function importCSV() {
    files.value = fileInput.value?.files

    if (files.value.length === 0) return

    let file = files.value[0] as File

    if (!file.name.endsWith(".csv") && file.type !== "text/csv") return

    let reader = new FileReader()
    reader.readAsText(file)
    reader.onload = () => {
        data.csv = reader.result
    }

}

</script>
