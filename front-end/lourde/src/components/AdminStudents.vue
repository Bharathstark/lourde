<template>
  <v-main>
    <v-toolbar flat>
      <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
      <AppBar v-model="toggle"></AppBar>
      <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
      <v-toolbar-title>
        Admin-Students
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        large
        icon
        v-on:click="(button = 'Save'), clearForm(), (studentDialog = true)"
      >
        <v-icon large>mdi-plus</v-icon>
      </v-btn>
    </v-toolbar>
    <v-container>
      <v-card flat>
        <v-tabs :key="sync" color="deep-purple accent-4" centered>
          <v-tab>12 Standard Results</v-tab>
          <v-tab>10 Standard Results</v-tab>
          <v-tab-item v-for="(students, i) in studentData" :key="i">
            <v-container>
              <v-row>
                <v-col
                  v-for="student in students"
                  :key="student.id"
                  cols="12"
                  xs="6"
                  sm="4"
                  md="3"
                  class="d-flex flex-column"
                >
                  <v-card class=" flex d-flex flex-column">
                    <div class=" mt-8 text-center">
                      <v-avatar size="130">
                        <v-img large :src="student.thumbnail"></v-img>
                      </v-avatar>
                    </div>

                    <v-list-item>
                      <v-list-item-content>
                        <v-list-item-title class="ma-2 text-center">{{
                          student.name
                        }}</v-list-item-title>
                        <v-list-item-subtitle class=" text-center">{{
                          student.description
                        }}</v-list-item-subtitle>
                      </v-list-item-content>
                    </v-list-item>
                    <v-card-actions>
                      <v-btn text color="primary" @click="editForm(student)">
                        Edit
                      </v-btn>
                      <v-btn text color="error" @click="deleteEntity(student)">
                        Delete
                      </v-btn>
                      <v-btn text color="green" @click="showMarkForm(student)">
                        Add Mark
                      </v-btn>
                    </v-card-actions>
                    <v-list-item
                      class="flex ml-3 mb-0"
                      v-for="mark in student.marks"
                      :key="mark.id"
                    >
                      <v-list-item-content>{{
                        mark.field
                      }}</v-list-item-content>
                      <v-list-item-content>
                        {{ mark.value }}
                      </v-list-item-content>
                      <v-list-item-icon>
                        <v-btn
                          color="error"
                          @click="deleteMark(student, mark)"
                          icon
                        >
                          <v-icon>mdi-delete</v-icon>
                        </v-btn>
                      </v-list-item-icon>
                    </v-list-item>
                  </v-card>
                </v-col>
              </v-row>
            </v-container>
          </v-tab-item>
        </v-tabs>
      </v-card>
    </v-container>
    <v-dialog v-model="studentDialog" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">Student</span>
        </v-card-title>
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-card-text>
            <v-text-field
              label="Student Name"
              v-model="name"
              outlined
              required
            ></v-text-field>
            <v-text-field
              label="Sub Text"
              outlined
              required
              v-model="description"
            ></v-text-field>
            <v-row>
              <v-col cols="12" md="6">
                <v-select
                  v-model="entity"
                  :items="items"
                  label="Class"
                  outlined
                ></v-select>
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="order"
                  outlined
                  type="number"
                  label="Order"
                ></v-text-field>
              </v-col>
            </v-row>

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
            :loading="loading"
            color="blue darken-1"
            text
            @click="button == 'Save' ? saveEntity() : updateEntity()"
          >
            {{ button }}
          </v-btn>
          <v-btn color="blue darken-1" text @click="studentDialog = false">
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="markDialog" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">{{ studentname }}</span>
        </v-card-title>
        <v-form ref="markform" v-model="markValid" lazy-validation>
          <v-card-text>
            <v-text-field
              label="Student Field"
              v-model="field"
              outlined
              required
            ></v-text-field>
            <v-text-field
              label="Field Value"
              outlined
              required
              v-model="value"
            ></v-text-field>
          </v-card-text>
        </v-form>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            :loading="loading"
            color="blue darken-1"
            text
            @click="saveMarkEntity(studentid)"
          >
            Save
          </v-btn>
          <v-btn
            color="blue darken-1"
            text
            @click="(markDialog = false), (this.sync = this.sync + 1)"
          >
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
export default {
  name: "AdminStudents",
  data: () => ({
    valid: true,
    sync: 11,
    loading: false,
    markvalid: true,
    studentData: [],
    items: ["10", "12"],
    ten: [],
    twelve: [],
    apiURL: "/api/v1/students",
    date: null,
    toggle: false,
    studentDialog: false,
    studentname: "",
    studentid: 0,
    markDialog: false,
    id: null,
    name: null,
    button: "Save",
    description: null,
    entity: "",
    order: 0,
    files: null,
    datemenu: false
  }),
  components: {
    AppBar,
    confirmDialog: () => import("@/components/ConfirmDialog.vue")
  },
  props: {
    bindinfo: Object
  },
  created() {
    this.fetchCreateData();
    this.sync = this.sync + 1;
  },
  methods: {
    fetchData(url) {
      return this.axios
        .get(url)
        .then(response => {
          return response.data["data"];
        })
        .catch(error => {
          console.log(error);
        });
    },
    async fetchCreateData() {
      this.studentData = [];
      await this.fetchData(this.apiURL + "?module=12").then(data =>
        this.studentData.push(data)
      );
      await this.fetchData(this.apiURL + "?module=10").then(data =>
        this.studentData.push(data)
      );
    },
    editForm(value) {
      this.id = value.id;
      this.name = value.name;
      this.description = value.description;
      this.button = "Update";
      this.studentDialog = true;
      this.entity = value.entity;
      this.order = value.order;
    },
    clearForm() {
      this.id = "";
      this.name = "";
      this.description = "";
      this.entity = "";
      this.order = 0;
    },
    delete(url) {
      this.$refs.confirm
        .open("Delete", "Are you sure?", { color: "red" })
        .then(confirm => {
          if (confirm) {
            this.axios.delete(url).then(response => {
              this.id = null;
              if (response.data.response_status.response_code == 200) {
                this.fetchCreateData();
              }
            });
          }
        });
    },
    deleteEntity(value) {
      this.delete(this.apiURL + "/" + value.id);
      this.sync = this.sync + 1;
    },
    deleteMark(student, mark) {
      this.delete(this.apiURL + "/" + student.id + "/marks/" + mark.id);
    },
    saveEntity() {
      this.$refs.confirm
        .open("Save", "Are you sure?", { color: "green" })
        .then(confirm => {
          if (confirm) {
            this.loading = true;
            var formData = new FormData();
            formData.append("name", this.name);
            formData.append("description", this.description);
            formData.append("files", this.files);
            formData.append("order", this.order);
            this.axios
              .post(this.apiURL + "?module=" + this.entity, formData, {
                headers: {
                  "Content-Type": "multipart/form-data"
                }
              })
              .then(response => {
                if (response.data.response_status.response_code == 200) {
                  this.loading = false;
                  this.$refs.form.reset();
                  this.studentDialog = false;
                  this.fetchCreateData();
                }
              });
          }
        });
    },
    showMarkForm(value) {
      this.studentid = value.id;
      this.studentname = value.name;
      this.markDialog = true;
    },
    saveMarkEntity(studentid) {
      this.$refs.confirm
        .open("Save", "Are you sure?", { color: "green" })
        .then(confirm => {
          if (confirm) {
            this.loading = true;
            this.axios
              .post(this.apiURL + "/" + studentid + "/marks", {
                field: this.field,
                value: this.value
              })
              .then(response => {
                if (response.data.response_status.response_code == 200) {
                  this.loading = false;
                  this.$refs.markform.reset();
                  this.fetchCreateData();
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
            this.loading = true;
            var formData = new FormData();
            formData.append("name", this.name);
            formData.append("description", this.description);
            formData.append("files", this.files);
            formData.append("order", this.order);
            formData.append("entity", this.entity);

            this.axios
              .put(this.apiURL + "/" + this.id, formData, {
                headers: {
                  "Content-Type": "multipart/form-data"
                }
              })
              .then(response => {
                if (response.data.response_status.response_code == 200) {
                  this.loading = false;
                  this.$refs.form.reset();
                  this.studentDialog = false;
                  this.fetchCreateData();
                }
              });
          }
        });
    }
  }
};
</script>

<style></style>
