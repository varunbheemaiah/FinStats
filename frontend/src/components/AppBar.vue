<template>
    <v-app-bar app elevation="0">
        <v-toolbar-title>{{ pageTitle }}</v-toolbar-title>
    </v-app-bar>
</template>

<script lang="ts" setup>
import { toRefs } from 'vue'
const props = defineProps({
    pageTitle: String,
})
const { pageTitle } = toRefs(props)
</script>

<style>
.v-app-bar {
    height: 67.5px;
}
</style>