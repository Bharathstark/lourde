<template>
  <v-dialog
    v-model="dialog"
    @keydown.esc="cancel"
    :fullscreen="$vuetify.breakpoint.xs || $vuetify.breakpoint.sm"
    :hide-overlay="$vuetify.breakpoint.xs || $vuetify.breakpoint.sm"
    :max-width="$vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 0 : 500"
  >
    <v-card>
      <v-toolbar>
        <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
        <v-toolbar-title>
          Notifications
        </v-toolbar-title>
        <v-spacer></v-spacer>

        <v-btn icon light @click="dialog = false">
          <v-icon color="grey darken-2">
            mdi-close
          </v-icon>
        </v-btn>
      </v-toolbar>
      <v-list class="overflow-y-auto" height="500">
        <v-list-item v-for="notification in notiData" :key="notification.id">
          <v-list-item-content class="wrap-text">
            {{ notification.name }}
          </v-list-item-content>
          <v-list-item-icon v-show="notification.url != ''">
            <v-btn icon :href="notification.url" target="_blank">
              <v-icon>mdi-open-in-new</v-icon>
            </v-btn>
          </v-list-item-icon>
        </v-list-item>
      </v-list>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "Notifications",
  data: () => ({
    dialog: false,
    notiURL: "/api/v1/notifications?module=notification",
    notiData: []
  }),
  methods: {
    open() {
      this.dialog = true;
      this.fetchData();
    },
    cancel() {
      this.dialog = false;
    },
    fetchData() {
      this.axios
        .get(this.notiURL)
        .then(response => {
          this.notiData = response.data["data"];
        })
        .catch(error => {
          console.log(error);
        });
    }
  },
  created() {
    this.fetchData();
  }
};
</script>

<style></style>
