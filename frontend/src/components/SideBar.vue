<template>
    <v-navigation-drawer app>
        <v-list-item title="FinStats" subtitle="The one stop shop for all finance statistics"></v-list-item>
        <v-divider></v-divider>
        <router-link to="/nasdaq-earnings">
            <v-list-item link title="NASDAQ Earnings Calendar"></v-list-item>
        </router-link>
        <router-link to="/greeting">
            <v-list-item link title="Greeting"></v-list-item>
        </router-link>
    </v-navigation-drawer>
</template>
  
<script setup lang="ts">
// Add any necessary script setup code here
</script>
  
<style scoped>
.v-list-item:first-child {
    margin: 0.75em 0em;
}
a {
    text-decoration: none;
    color: white;
    opacity: 0.7;
}
</style>
  