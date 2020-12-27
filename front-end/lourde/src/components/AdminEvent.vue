<template>
  <v-main>
    <v-card flat height="100%">
      <v-toolbar color="white" flat>
        <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
        <AppBar v-model="toggle"></AppBar>
        <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
        <v-toolbar-title class="grey--text text--darken-4">
          {{ bindinfo.module }}
        </v-toolbar-title>
        <v-spacer></v-spacer>

        <v-btn large icon v-on:click="saveForm">
          <v-icon large>mdi-plus</v-icon>
        </v-btn>
      </v-toolbar>
      <v-container height="100%">
        <v-row fluid>
          <v-col
            v-for="i in eventData"
            :key="i.id"
            cols="6"
            md="3"
            sm="4"
            xs="6"
          >
            <v-hover v-slot:default="{ hover }">
              <v-card
                v-on:click="
                  $emit('showEvents', { id: i.id, module: bindinfo.module })
                "
                :height="
                  $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 180 : 300
                "
                :width="
                  $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 150 : 300
                "
              >
                <v-img class="event" height="75%" :src="i.thumbnail"> </v-img>
                <v-card-title class="text-lg-h6 text-subtitle-2">
                  {{ i.name }}
                </v-card-title>

                <v-fade-transition>
                  <v-overlay v-show="hover" absolute>
                    <v-btn icon @click.stop="editForm(i)" class="ma-3">
                      <v-icon large elevation="6" color="primary">
                        mdi-pencil
                      </v-icon>
                    </v-btn>
                    <v-btn icon @click.stop="deleteEntity(i)" class="ma-3">
                      <v-icon large elevation="6" color="error">
                        mdi-delete
                      </v-icon>
                    </v-btn>
                  </v-overlay>
                </v-fade-transition>
              </v-card>
            </v-hover>
          </v-col>
        </v-row>
      </v-container>
      <v-dialog v-model="dialog" persistent max-width="600px">
        <v-card>
          <v-card-title>
            <span class="headline">{{ bindinfo.module }} </span>
          </v-card-title>
          <v-form ref="form" v-model="valid" lazy-validation>
            <v-card-text>
              <v-text-field
                label="Name"
                v-model="name"
                outlined
                required
              ></v-text-field>

              <v-text-field
                label="Description"
                outlined
                required
                v-model="description"
              ></v-text-field>
              <v-menu
                v-model="datemenu"
                :close-on-content-click="false"
                :nudge-right="40"
                transition="scale-transition"
                offset-y
                min-width="290px"
              >
                <template v-slot:activator="{ on, attrs }">
                  <v-text-field
                    v-model="date"
                    outlined
                    label="Event Date"
                    prepend-icon="mdi-calendar"
                    readonly
                    v-bind="attrs"
                    v-on="on"
                  ></v-text-field>
                </template>
                <v-date-picker
                  v-model="date"
                  @input="datemenu = false"
                ></v-date-picker>
              </v-menu>

              <v-file-input
                counter
                show-size
                small-chips
                outlined
                truncate-length="5"
                label="Upload File"
                required
                accept="image/*"
                v-model="files"
              ></v-file-input>
            </v-card-text>
          </v-form>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="button == 'Save' ? saveEntity() : updateEntity()"
            >
              {{ button }}
            </v-btn>
            <v-btn color="blue darken-1" text @click="dialog = false">
              Close
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-card>
    <confirmDialog ref="confirm" />
  </v-main>
</template>

<script>
import AppBar from "@/components/AppBar.vue";
import confirmDialog from "@/components/ConfirmDialog.vue";
export default {
  name: "AdminEvent",
  data: () => ({
    valid: true,
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
    datemenu: false
  }),
  props: {
    bindinfo: Object
  },
  components: {
    AppBar,
    confirmDialog
  },
  methods: {
    fetchData() {
      this.axios
        .get(this.apiURL + "?module=" + this.bindinfo.module.toLowerCase())
        .then(response => {
          this.eventData = response.data["data"];
        })
        .catch(error => {
          console.log(error);
        });
    },
    saveEntity() {
      this.$refs.confirm
        .open("Save", "Are you sure?", { color: "green" })
        .then(confirm => {
          if (confirm) {
            var formData = new FormData();
            formData.append("name", this.name);
            formData.append("description", this.description);
            formData.append("files", this.files);
            formData.append("date", new Date(this.date).getTime());
            this.axios
              .post(
                this.apiURL + "?module=" + this.bindinfo.module.toLowerCase(),
                formData,
                {
                  headers: {
                    "Content-Type": "multipart/form-data"
                  }
                }
              )
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
    updateEntity() {
      this.$refs.confirm
        .open("Update", "Are you sure?", { color: "blue" })
        .then(confirm => {
          if (confirm) {
            var formData = new FormData();
            formData.append("name", this.name);
            formData.append("description", this.description);
            formData.append("files", this.files);
            formData.append("date", new Date(this.date).getTime());
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
    deleteEntity(value) {
      this.$refs.confirm
        .open("Delete", "Are you sure?", { color: "red" })
        .then(confirm => {
          if (confirm) {
            this.axios.delete(this.apiURL + "/" + value.id).then(response => {
              this.id = null;
              if (response.data.response_status.response_code == 200) {
                this.fetchData();
              }
            });
          }
        });
    },
    editForm(value) {
      this.id = value.id;
      (this.name = value.name), (this.description = value.description);
      this.date = new Date(value.date).toISOString().substring(0, 10);
      this.button = "Update";
      this.dialog = true;
    },
    saveForm() {
      this.dialog = true;
      this.$refs.form.reset();
      this.button = "Save";
    }
  },
  created() {
    this.fetchData();
  }
};
</script>

<style></style>
