<template>
  <el-row>
  </el-row>
  <el-table :data="tableData" style="width: 100%; margin-top: 10px; margin-bottom: 10px; height: 85%">
    <el-table-column label="用户ID" prop="id" />
    <el-table-column label="用户名" prop="username" />
    <el-table-column label="昵称" prop="full_name" />
    <el-table-column label="简介" prop="bio" />
    <el-table-column label="角色码" prop="role" />
    <el-table-column label="头像" prop="avatar_url" />
    <el-table-column align="center" label="操作">
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)">Edit</el-button>
        <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">Delete</el-button>
      </template>
    </el-table-column>
  </el-table>
  <div class="pagination">
    <el-pagination
        v-model:current-page="currentPage4"
        v-model:page-size="pageSize4"
        :page-sizes="[5, 10, 20, 40]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="400"
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
    />
  </div>
</template>

<script lang="ts" setup>
import { Timer } from '@element-plus/icons-vue'
import { ref } from "vue";

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

const currentPage4 = ref(4)
const pageSize4 = ref(100)
const handleSizeChange = (val: number) => {
  console.log(`${val} items per page`)
}
const handleCurrentChange = (val: number) => {
  console.log(`current page: ${val}`)
}

const tableData: User[] = [
  {
    id: 1,
    username: 'user1',
    full_name: 'User One',
    bio: 'Bio for User One',
    role: 1,
    avatar_url: 'https://example.com/avatar1.jpg'
  },
  {
    id: 2,
    username: 'user2',
    full_name: 'User Two',
    bio: 'Bio for User Two',
    role: 2,
    avatar_url: 'https://example.com/avatar2.jpg'
  },
  {
    id: 3,
    username: 'user3',
    full_name: 'User Three',
    bio: 'Bio for User Three',
    role: 1,
    avatar_url: null
  },
  {
    id: 4,
    username: 'user4',
    full_name: 'User Four',
    bio: 'Bio for User Four',
    role: 3,
    avatar_url: 'https://example.com/avatar4.jpg'
  }
];

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