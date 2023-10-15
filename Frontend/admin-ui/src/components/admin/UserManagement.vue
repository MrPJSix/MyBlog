<template>
  <el-row>
  </el-row>
  <el-table :data="tableData" style="width: 100%; margin-top: 10px; margin-bottom: 10px; height: 85%">
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
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)">Edit</el-button>
        <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">Delete</el-button>
      </template>
    </el-table-column>
  </el-table>
  <div class="pagination">
    <el-pagination
        v-model:current-page="pageNum"
        v-model:page-size="pageSize"
        :page-sizes="[1, 2, 5, 10, 20, 40]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="usersCount"
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
    />
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted, nextTick } from "vue";
  import axios from "axios";

  interface User {
    id: number;
    username: string;
    full_name: string;
    bio: string;
    role: number;
    avatar_url?: string | null;
  }

  const handleEdit = (index: number, row: User) => {
    console.log(index, row)
  }
  const handleDelete = (index: number, row: User) => {
    console.log(index, row)
  }

  const usersCount = ref(4);
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
      const response = await axios.get(`http://124.220.25.230:9000/admin/users`, {
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

  onMounted(() => {
    fetchData(); // 初始加载数据
  });

  onMounted(async () => {
    try response = await 
  })
  const tableData = ref<User[]>([]);

</script>

<style scoped>
.user-management{
  align-items: center;
  justify-content: space-between;
  text-align: center;
}
.pagination {
  height: 10%;
  width: 100%;
  display: flex;
  justify-content: center;
}
</style>