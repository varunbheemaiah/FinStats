<script lang="ts" setup>
import SideBar from './components/SideBar.vue'
import AppBar from './components/AppBar.vue'
import { ref } from 'vue';
import router from './router';

const pageTitle = ref(document.title);

router.beforeEach((to, from, next) => {
  pageTitle.value = String(to.name) || 'Page';
  next();
});
</script>

<template>
  <v-app>
    <SideBar />
    <v-main>
      <AppBar :pageTitle="pageTitle" />
      <router-view/>
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
