<template>
  <div class="user-management">
    <el-row>
      <div style="height: 10vh; width: 100%; display: flex; align-items: center;" >
        <input  v-model="searchedId"
                placeholder="Type ID to search."
                class="input-search"
                @keyup.enter="searchUserByID"
        >
        <el-button :icon="Search"
                   circle style="margin-left: 10px"
                   @click="searchUserByID"
        />
      </div>
    </el-row>
    <el-row>
      <el-table :data="tableData" height="68vh" max-height="68vh" style="width: 100%; margin-top: 1vh; margin-bottom: 1vh;">
        <el-table-column label="用户ID" prop="id" />
        <el-table-column label="用户名" prop="username" />
        <el-table-column label="昵称" prop="full_name" />
        <el-table-column label="简介" prop="bio" />
        <el-table-column label="角色码" prop="role" />
        <el-table-column label="头像" prop="avatar_url">
          <template #default="scope">
            <el-avatar :size="36" :src="scope.row.avatar_url" />
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="deleteUserByID(scope.row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-row>
    <el-row>
      <div class="pagination" style="height: 10vh;">
        <el-pagination
            v-model:current-page="pageNum"
            v-model:page-size="pageSize"
            :page-sizes="[5, 10, 20, 40]"
            layout="total, sizes, prev, pager, next, jumper"
            :total="totalUsers"
            background
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
        />
      </div>
    </el-row>
    <el-row>
      <el-dialog v-model="dialogVisible" title="修改用户信息" width="500px">
        <el-form :model="editedUserData" style="display: flex; flex-direction: column; align-items: center">
          <el-form-item label="用户ID" prop="id" label-width="100px">
            <el-input v-model="editedUserData.id" disabled style="width:300px"/>
          </el-form-item>
          <el-form-item label="用户名" prop="username" label-width="100px">
            <el-input v-model="editedUserData.username" disabled style="width:300px"/>
          </el-form-item>
          <el-form-item label="昵称" prop="full_name" label-width="100px">
            <el-input v-model="editedUserData.full_name" style="width:300px"/>
          </el-form-item>
          <el-form-item label="简介" prop="bio" label-width="100px">
            <el-input v-model="editedUserData.bio" style="width:300px"/>
          </el-form-item>
          <el-form-item label="角色码" prop="role" label-width="100px">
            <el-input v-model="editedUserData.role" disabled style="width:300px"/>
          </el-form-item>
          <el-form-item label="头像" prop="avatar_url" label-width="100px">
            <el-input v-model="editedUserData.avatar_url" disabled style="width:300px"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="editUser">
              确认
            </el-button>
          </span>
        </template>
      </el-dialog>
    </el-row>
  </div>

</template>

<script setup>
  import { Search } from "@element-plus/icons-vue"
  import {ref, onMounted, nextTick, reactive} from "vue";
  import axios from "axios";
  import { ElMessage, ElMessageBox } from "element-plus";
  import baseURL from "../../api/api.config";

  const tableData = ref([]);

  /* ---------- 分页查询-Start ---------- */

  const totalUsers = ref(0);
  const pageNum = ref(1);
  const pageSize = ref(5);
  const handleSizeChange = (val) => {
    console.log(`${val} items per page`);
    pageSize.value = val; // 在更改时更新页面大
    nextTick(() => {
      fetchData();
    });
  }
  const handleCurrentChange = (val) => {
    console.log(`current page: ${val}`);
    pageNum.value = val; // 在更改时更新当前页面
    nextTick(() => {
      fetchData();
    });
  }
  const fetchData = async () => {
    try {
      const response = await axios.get(`${baseURL}/users`, {
        params: {
          pagesize: pageSize.value,
          pagenum: pageNum.value,
        },
      });
      if (response.data.status === 200) {
        tableData.value = response.data.data;
      }
    } catch (error) {
      console.error('An error occurred:', error);
    }
  }
  const fetchTotal = async () => {
    try {
      const response = await axios.get(`${baseURL}/users/all`);
      if (response.data.status === 200) {
        totalUsers.value = response.data.data;
      }
    } catch (error) {
      console.error('An error occurred:', error);
    }
  }
  onMounted(() => {
    fetchTotal();
    fetchData(); // 初始加载数据
  });
  /* ---------- 分页查询-End ---------- */

  /* ---------- 按ID查询-Start ---------- */
  const searchedId = ref("");
  const searchUserByID = async () => {
    try {
      const response = await axios.get(`${baseURL}/user/${searchedId.value}`);
      if (response.data.status === 200) {
        ElMessage({
          message: '查询成功',
          type: 'success',
        });
        tableData.value = [response.data.data];
      } else {
        ElMessage({
          message: response.data.message,
          type: 'error',
        });
      }
    } catch (error) {
      ElMessage({
        message: 'An error occurred while trying to log in.',
        type: 'warning',
      })
      console.error('An error occurred:', error);
    }
  }
  /* ---------- 按ID查询-End ---------- */

  /* ---------- 按ID删除-Start ---------- */
  const deleteUserByID = async (userID) => {
    ElMessageBox.confirm(
        '此操作会永久删除数据，是否继续？',
        '警告',
        {
          confirmButtonText: '确认',
          cancelButtonText: '取消',
          type: 'warning',
        }
    ).then(async () => {
        try {
          const response = await axios.delete(`${baseURL}/user/${userID}`);
          if (response.data.status === 200) {
            ElMessage({
              message: '删除成功',
              type: 'success',
            });
            totalUsers.value--;
            await nextTick(() => {
              fetchData();
            });
          } else {
            ElMessage({
              message: response.data.message,
              type: 'error',
            });
          }
        } catch (error) {
          ElMessage({
            message: 'An error occurred while trying to log in.',
            type: 'warning',
          })
          console.error('An error occurred:', error);
        }
      }).catch(() => {
          ElMessage({
            type: 'info',
            message: '删除取消',
          })
        })
  }
  /* ---------- 按ID删除-End ---------- */

  /* ---------- 编辑用户-End ---------- */
  const editedUserData = reactive({
    id: '',
    username: '',
    full_name: '',
    bio: '',
    role: '',
    avatar_url:''
  });
  const dialogVisible = ref(false);
  const handleEdit = (index, row) => {
    dialogVisible.value = true;
    editedUserData.id = row.id;
    editedUserData.username = row.username;
    editedUserData.full_name = row.full_name;
    editedUserData.bio = row.bio;
    editedUserData.role = row.role;
    editedUserData.avatar_url = row.avatar_url;
    console.log(dialogVisible.value);
    console.log("Editing ", index, row);
  }
  const editUser = async () => {
    ElMessageBox.confirm(
        '此操作会永久修改数据，是否继续？',
        '警告',
        {
          confirmButtonText: '确认',
          cancelButtonText: '取消',
          type: 'warning',
        }
    ).then(async () => {
      try {
        const response = await axios.put(`${baseURL}/user/${editedUserData.id}`, {
          full_name: editedUserData.full_name,
          bio: editedUserData.bio
        });
        if (response.data.status === 200) {
          ElMessage({
            message: '更新成功',
            type: 'success',
          });
          totalUsers.value--;
          dialogVisible.value = false;
          await nextTick(() => {
            fetchData();
          });
        } else {
          ElMessage({
            message: response.data.message,
            type: 'error',
          });
        }
      } catch (error) {
        ElMessage({
          message: 'An error occurred while trying to edit.',
          type: 'warning',
        })
        console.error('An error occurred:', error);
      }
    }).catch(() => {
      ElMessage({
        type: 'info',
        message: '更新取消',
      })
    })
  }
  /* ---------- 编辑用户-End ---------- */
</script>

<style scoped>
  .user-management{
    margin: 0;
    padding-left: 20px;
    padding-right: 20px;
    height: 90vh;
    //align-items: center;
    //justify-content: space-between;
    //text-align: center;
  }

  .input-search {
    margin-left: 10px;
    height: 40%;
    border-radius: 10px;
    border-color: #c6d0dc;
    border-style: solid;
    text-align: center;
  }

  .pagination {
    width: 100%;
    display: flex;
    justify-content: center;
  }
</style>