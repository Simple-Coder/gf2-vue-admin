<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="参数名称">
        <el-input v-model="queryParams.configName" placeholder="请输入参数名称" style="width: 240px" clearable size="small" />
      </el-form-item>
      <el-form-item label="参数键名">
        <el-input v-model="queryParams.configKey" placeholder="请输入参数键名" style="width: 240px" clearable size="small" />
      </el-form-item>
      <el-form-item label="系统内置">
        <el-select placeholder="系统内置">
          <el-option />
        </el-select>
      </el-form-item>
      <el-form-item label="创建时间">
        <el-date-picker v-model="dateRange" style="width: 240px" range-separator="-" size="small" type="daterange" value-format="yyyy-MM-dd" start-placeholder="开始日期" end-placeholder="结束日期" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" size="mini" icon="el-icon-plus">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" size="mini" icon="el-icon-edit">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" size="mini" icon="el-icon-delete">删除</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="warning" size="mini" icon="el-icon-download">导出</el-button>
      </el-col>
    </el-row>

    <el-table v-loading="listLoading" :data="list">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="参数主键" align="center" prop="configId" />
      <el-table-column label="参数名称" align="center" prop="configName" />
      <el-table-column label="参数键名" align="center" prop="configKey" />
      <el-table-column label="参数键值" align="center" prop="configValue" />
      <el-table-column label="系统内置" align="center" prop="configType" />
      <el-table-column label="备注" align="center" prop="remark" />
      <el-table-column label="创建时间" align="center" prop="createdAt" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)">修改</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getSysConfigList } from '@/api/system/sysconfig'

export default {
  name: 'SysParam',
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true,
      // 日期范围
      dateRange: [],
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        configName: undefined,
        configKey: undefined,
        configType: undefined
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getSysConfigList().then(response => {
        this.list = response.data.list
        this.listLoading = false
      })
    }
  }
}
</script>
<style lang="scss" scoped>
.mb8 {
  margin-bottom: 8px;
}
</style>
