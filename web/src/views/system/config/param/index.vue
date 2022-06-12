<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
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
