<template>
    <v-navigation-drawer app>
        <router-link to="/">
            <v-list-item class="no-hover">
                <template #prepend>
                    <v-list-item-avatar class="square-avatar">
                        <img src="../assets/images/stock-market-no-base.png" alt="FinStats Image">
                    </v-list-item-avatar>
                </template>
                <v-list-item-content>
                    <v-list-item-title>FinStats</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </router-link>
        <v-divider></v-divider>
        <router-link to="/nasdaq-earnings">
            <v-list-item link title="NASDAQ Earnings Calendar"></v-list-item>
        </router-link>
    </v-navigation-drawer>
</template>
  
<script setup lang="ts">
// Add any necessary script setup code here
</script>
  
<style scoped>
.no-hover:hover,
.no-hover:focus {
    background-color: transparent !important;
}

.v-list-item:first-child {
    margin: 0.5em 0em;
}

a {
    text-decoration: none;
    color: white;
    opacity: 0.7;
}

.square-avatar img {
    width: 36px;
    /* border-radius: 0; */
}
</style>
  