<template>
  <v-app>
    <router-view />
  </v-app>
</template>

<script>
export default {
  name: "App",

  data: () => ({
    //
  }),
  components: {},
  created() {
    window.addEventListener("beforeinstallprompt", e => {
      e.prompt();
    });
  }
};
</script>
<style>
@font-face {
  font-family: "Coronet";
  font-style: normal;
  font-weight: 100;
  src: url("./assets/fonts/Coronet.woff");
}
</style>
