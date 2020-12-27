<template>
  <v-main>
    <v-toolbar color="white" flat>
      <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
      <AppBar v-model="toggle"></AppBar>

      <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
      <v-toolbar-title class="grey--text text--darken-4">
        Admin-Principal
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn large icon v-on:click="dialog = true">
        <v-icon large>mdi-plus</v-icon>
      </v-btn>
    </v-toolbar>
    <section>
      <v-card dark flat>
        <v-img class="vimage" height="600" :src="this.url">
          <v-row>
            <v-col cols="6" md="5" class="ml-auto">
              <v-card-text
                :style="
                  $vuetify.breakpoint.xs || $vuetify.breakpoint.sm
                    ? styled2
                    : styled
                "
                class="text-center"
              >
                Principal Address
              </v-card-text>
              <v-card-text
                :style="
                  $vuetify.breakpoint.xs || $vuetify.breakpoint.sm
                    ? styled2
                    : styled
                "
              >
                {{ this.description }}
              </v-card-text>
              <v-row>
                <v-col md="5" cols="6" class="ml-auto">
                  <v-card-text
                    :style="
                      $vuetify.breakpoint.xs || $vuetify.breakpoint.sm
                        ? styled2
                        : styled
                    "
                  >
                    - {{ this.name }}
                  </v-card-text>
                </v-col>
              </v-row>
            </v-col>
          </v-row>
        </v-img>
      </v-card>
    </section>
    <v-dialog v-model="dialog" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">Principal </span>
        </v-card-title>
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-card-text>
            <v-text-field
              label="Name"
              v-model="name"
              outlined
              required
            ></v-text-field>

            <v-textarea
              label="Description"
              auto-grow
              outlined
              required
              v-model="description"
            ></v-textarea>

            <v-row>
              <v-col cols="6">
                <v-file-input
                  counter
                  show-size
                  small-chips
                  outlined
                  truncate-length="5"
                  label="Upload Profile Photo"
                  required
                  accept="image/*"
                  v-model="files"
                ></v-file-input>
              </v-col>
              <v-col cols="6">
                <v-switch
                  v-model="date"
                  value="1"
                  label="show profile"
                ></v-switch>
              </v-col>
            </v-row>
          </v-card-text>
        </v-form>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="editForm">
            Update
          </v-btn>
          <v-btn color="blue darken-1" text @click="dialog = false">
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <confirmDialog ref="confirm" />
  </v-main>
</template>

<script>
import AppBar from "@/components/AppBar.vue";
import confirmDialog from "@/components/ConfirmDialog.vue";
export default {
  name: "AdminPrincipal",
  data: () => ({
    valid: true,
    styled: "font-family:'Coronet';font-size:35px;line-height: 1.1",
    styled2: "font-family:'Coronet';font-size:19px;line-height: 1.0",
    eventData: null,
    apiURL: "/api/v1/events",
    date: null,
    toggle: false,
    dialog: false,
    id: null,
    name: null,
    button: "Save",
    description: null,
    files: null,
    url: ""
  }),
  components: {
    AppBar,
    confirmDialog
  },
  created() {
    this.fetchData();
  },
  watch: {
    eventData: function(newValue) {
      this.id = newValue.id;
      this.name = newValue.name;
      this.description = newValue.description;
      this.date = newValue.date == 1 ? "1" : "2";
      this.url = newValue.thumbnail;
    }
  },
  methods: {
    fetchData() {
      this.axios
        .get(this.apiURL + "?module=principal")
        .then(response => {
          this.eventData = response.data["data"][0];
        })
        .catch(error => {
          console.log(error);
        });
    },
    updateEntity() {
      this.$refs.confirm
        .open("Update", "Are you sure?", { color: "blue" })
        .then(confirm => {
          if (confirm) {
            this.date = this.date != "1" ? "2" : "1";
            var formData = new FormData();
            formData.append("name", this.name);
            formData.append("description", this.description);
            formData.append("files", this.files);
            formData.append("date", this.date);
            this.axios
              .put(this.apiURL + "/" + this.id, formData, {
                headers: {
                  "Content-Type": "multipart/form-data"
                }
              })
              .then(response => {
                if (response.data.response_status.response_code == 200) {
                  this.$refs.form.reset();
                  this.dialog = false;
                  this.fetchData();
                }
              });
          }
        });
    },
    editForm() {
      this.updateEntity();
    }
  }
};
</script>

<style scoped>
.vimage .v-image__image {
  background-size: 100% 100%;
}
</style>
