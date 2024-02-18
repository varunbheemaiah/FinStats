<script lang="ts" setup>
import {reactive} from 'vue'

const data = reactive({
  inputType: 'range',
})

function submit() {
  console.log(data.inputType)
}

</script>

<template>
  <main>
    <v-radio-group inline label="Will you enter a date range or multiple dates?" v-model="data.inputType">
      <v-radio label="Date Range" value="range"></v-radio>
      <v-radio label="Multiple Dates" value="dates"></v-radio>
    </v-radio-group>
    <v-date-picker></v-date-picker>
    <v-btn variant="tonal" @click="submit">
      Submit
    </v-btn>
  </main>
</template>

<style>
main {
  padding: 1em 1em;
}

.v-radio-group > .v-input__control > .v-label {
  margin-inline-start: 0 !important;
}

.v-radio input[type="radio"] {
  opacity: 1 !important;
}

.v-radio {
  padding: 1em 0em;
}

.v-radio .v-label {
  margin-left: 1em;
}
</style>