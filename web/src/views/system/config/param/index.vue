<template>
  <div class="app-container">
    <el-form :inline="true">
      <el-form-item label="参数名称">
        <el-input placeholder="请输入参数名称" />
      </el-form-item>
      <el-form-item label="参数键名">
        <el-input placeholder="请输入参数键名" />
      </el-form-item>
      <el-form-item label="系统内置">
        <el-select placeholder="系统内置">
          <el-option />
        </el-select>
      </el-form-item>
      <el-form-item label="创建时间">
        <el-date-picker range-separator="-" size="small" type="daterange" value-format="yyyy-MM-dd" start-placeholde="开始日期" end-placeholde="结束日期" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini">搜索</el-button>
        <el-button  icon="el-icon-refresh" size="mini">重置</el-button>
      </el-form-item>
    </el-form>

    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">
          {{ scope.$index }}
        </template>
      </el-table-column>
      <el-table-column label="参数名">
        <template slot-scope="scope">
          {{ scope.row.configName }}
        </template>
      </el-table-column>
      <el-table-column label="参数值" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.configKey }}</span>
        </template>
      </el-table-column>
      <el-table-column label="参数键" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.configValue }}
        </template>
      </el-table-column>
      <!-- <el-table-column class-name="status-col" label="Status" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column> -->
      <el-table-column align="center" prop="remark" label="备注" width="200">
        <template slot-scope="scope">
          <!-- <i class="el-icon-time" /> -->
          <span>{{ scope.row.remark }}</span>
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
      listLoading: true
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
