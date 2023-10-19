<template>
  <div class="category-management">
    <el-row>
      <el-col :span="8">
        <div style="height: 10vh; display: flex; align-items: center;" >
          <input  v-model="searchedId"
                  placeholder="Type ID to search."
                  class="input-search"
                  @keyup.enter="searchCategoryByID"
          >
          <el-button :icon="Search"
                     circle style="margin-left: 10px"
                     @click="searchCategoryByID"
          />
        </div>
      </el-col>
      <el-col :span="8" :offset="8">
        <div style="height: 10vh; display: flex; align-items: center; float: right; padding-right: 50px">
          <el-button color="skyblue"><p style="color: white; font-weight: bold;">创建</p></el-button>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-table :data="tableData" height="68vh" max-height="68vh" style="width: 100%; margin-top: 1vh; margin-bottom: 1vh;">
        <el-table-column label="分类ID" prop="id" />
        <el-table-column label="分类名" prop="name" />
        <el-table-column align="center" label="操作">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="deleteCategoryByID(scope.row.id)">删除</el-button>
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
            :total="totalCategorys"
            background
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
        />
      </div>
    </el-row>
    <el-row>
      <el-dialog v-model="dialogVisible" title="修改分类信息" width="500px">
        <el-form :model="editedCategoryData" style="display: flex; flex-direction: column; align-items: center">
          <el-form-item label="分类ID" prop="id" label-width="100px">
            <el-input v-model="editedCategoryData.id" disabled style="width:300px"/>
          </el-form-item>
          <el-form-item label="分类名" prop="name" label-width="100px">
            <el-input v-model="editedCategoryData.username" style="width:300px"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="editCategory">
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

const totalCategorys = ref(0);
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
      totalCategorys.value = response.data.data;
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
const searchCategoryByID = async () => {
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
const deleteCategoryByID = async (userID) => {
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
        totalCategorys.value--;
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
const editedCategoryData = reactive({
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
  editedCategoryData.id = row.id;
  editedCategoryData.username = row.username;
  editedCategoryData.full_name = row.full_name;
  editedCategoryData.bio = row.bio;
  editedCategoryData.role = row.role;
  editedCategoryData.avatar_url = row.avatar_url;
  console.log(dialogVisible.value);
  console.log("Editing ", index, row);
}
const editCategory = async () => {
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
      const response = await axios.put(`${baseURL}/user/${editedCategoryData.id}`, {
        full_name: editedCategoryData.full_name,
        bio: editedCategoryData.bio
      });
      if (response.data.status === 200) {
        ElMessage({
          message: '更新成功',
          type: 'success',
        });
        totalCategorys.value--;
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
.category-management{
  margin: 0;
  padding-left: 20px;
  padding-right: 20px;
  height: 90vh;
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