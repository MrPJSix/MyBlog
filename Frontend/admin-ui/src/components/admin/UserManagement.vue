<template>
  <div class="user-management">
    <div style="height: 10%; width: 100%; display: flex; align-items: center;" >
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
    <el-table :data="tableData" style="width: 100%; margin-top: 10px; margin-bottom: 10px; height: 80%">
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
    <div class="pagination">
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
  </div>

</template>

<script lang="ts" setup>
  import { Search } from "@element-plus/icons-vue"
  import { ref, onMounted, nextTick } from "vue";
  import axios from "axios";
  import { ElMessage, ElMessageBox } from "element-plus";
  import baseURL from "../../api/api.config";

  interface User {
    id: number;
    username: string;
    full_name: string;
    bio: string;
    role: number;
    avatar_url?: string | null;
  }
  const tableData = ref<User[]>([]);

  /* ---------- 分页查询-Start ---------- */
  const handleEdit = (index: number, row: User) => {
    console.log(index, row)
  }

  const totalUsers = ref(0);
  const pageNum = ref(1);
  const pageSize = ref(5);
  const handleSizeChange = (val: number) => {
    console.log(`${val} items per page`);
    pageSize.value = val; // 在更改时更新页面大
    nextTick(() => {
      fetchData();
    });
  }
  const handleCurrentChange = (val: number) => {
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
  const deleteUserByID = async (userID: number) => {
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

</script>

<style scoped>
  .user-management{
    height: 100%;
    align-items: center;
    justify-content: space-between;
    text-align: center;
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
    height: 10%;
    width: 100%;
    display: flex;
    justify-content: center;
  }
</style>