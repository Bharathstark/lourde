<template>
  <v-dialog
    v-model="dialog"
    fullscreen
    hide-overlay
    transition="dialog-bottom-transition"
    @keydown.esc="cancel"
  >
    <v-card flat class="overflow-y-auto">
      <v-toolbar flat>
        <v-btn icon @click="dialog = false">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-toolbar-title>{{ data.name }}</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn text @click="dialog = false">
            Close
          </v-btn>
        </v-toolbar-items>
      </v-toolbar>
      <v-container fluid v-for="header in data.headers" :key="header.heading">
        <v-card-title
          class="justify-center text-md-body-1 text-sm-body-1 text-body-1 text-uppercase"
          >{{ header.heading }}</v-card-title
        >
        <v-row>
          <v-col
            v-for="i in header.images"
            :key="i.id"
            cols="6"
            md="3"
            sm="4"
            xs="6"
          >
            <v-img
              class="vimage"
              contain
              aspect-ratio="1"
              :src="'/img' + i"
            ></v-img>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "dialog",
  data: () => ({
    dialog: false,
    data: {}
  }),
  methods: {
    open(data) {
      this.data = data;
      this.dialog = true;
    }
  }
};
</script>
<style>
.vimage .v-image__image {
  background-size: 100% 100%;
}
</style>
