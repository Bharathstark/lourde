<template>
  <v-app>
    <confirmDialog ref="confirm" />
    <router-view :key="$route.fullPath" />
  </v-app>
</template>

<script>
export default {
  name: "App",
  components: {
    confirmDialog: () => import("@/components/ConfirmDialog.vue")
  },
  data: () => ({
    //
  }),
  created() {
    window.addEventListener("beforeinstallprompt", e => {
      e.preventDefault();
      this.$refs.confirm
        .open(
          "Install",
          "Do You Want to install this Application? No Additional Space Required",
          {
            color: "green"
          }
        )
        .then(confirm => {
          if (confirm) {
            e.prompt();
          }
        });
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
