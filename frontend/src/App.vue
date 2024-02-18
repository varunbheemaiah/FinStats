<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
import NasdaqEarningsCalendar from './components/NasdaqEarningsCalendar.vue';
import SideBar from './components/SideBar.vue'
</script>

<template>
  <v-app>
    <SideBar />
    <v-main>
      <NasdaqEarningsCalendar />
    </v-main>
  </v-app>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
