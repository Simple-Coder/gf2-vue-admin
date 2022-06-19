<template>
  <div class="app-container">
    <!--搜索框区域-->
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="参数名称">
        <el-input
          v-model="queryParams.configName"
          placeholder="请输入参数名称"
          style="width: 240px"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="参数键名">
        <el-input
          v-model="queryParams.configKey"
          placeholder="请输入参数键名"
          style="width: 240px"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="系统内置">
        <el-select v-model="queryParams.configType" placeholder="系统内置" clearable size="small">
          <!-- <el-option /> -->
        </el-select>
      </el-form-item>
      <el-form-item label="创建时间">
        <el-date-picker
          v-model="queryParams.dateRange"
          style="width: 240px"
          range-separator="-"
          size="small"
          type="daterange"
          value-format="yyyy-MM-dd"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
        />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <!--按钮区域-->
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" size="mini" icon="el-icon-plus" @click="handleAdd">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" size="mini" icon="el-icon-edit" :disabled="single" @click="handleUpdate">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" size="mini" icon="el-icon-delete" :disabled="multiple" @click="handleDelete">删除</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="warning" size="mini" icon="el-icon-download" @click="handleExport">导出</el-button>
      </el-col>
    </el-row>
    <!--表格数据区域-->
    <el-table v-loading="loading" :data="configList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="参数主键" align="center" prop="configId" />
      <el-table-column label="参数名称" align="center" prop="configName" />
      <el-table-column label="参数键名" align="center" prop="configKey" :show-overflow-tooltip="true" />
      <el-table-column label="参数键值" align="center" prop="configValue" :show-overflow-tooltip="true" />
      <el-table-column label="系统内置" align="center" prop="configType" />
      <el-table-column label="备注" align="center" prop="remark" :show-overflow-tooltip="true" />
      <el-table-column label="创建时间" align="center" prop="createdAt" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)">修改</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 添加或修改参数配置对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="500px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <!--      <el-form ref="form" :model="form" label-width="80px">-->
        <el-form-item label="参数名称" prop="configName">
          <el-input v-model="form.configName" placeholder="请输入参数名称" />
        </el-form-item>
        <el-form-item label="参数键名" prop="configKey">
          <el-input v-model="form.configKey" placeholder="请输入参数键名" />
        </el-form-item>
        <el-form-item label="参数键值" prop="configValue">
          <el-input v-model="form.configValue" placeholder="请输入参数键值" />
        </el-form-item>
        <el-form-item label="系统内置" prop="configType">
          <el-radio-group v-model="form.configType">
            <el-radio v-for="dict in typeOptions.values" :key="dict.key" :label="dict.key">{{ dict.value }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" placeholder="请输入内容" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getSysConfigList, addConfig, updateConfig, getConfig, delConfig, exportConfig } from '@/api/system/sysconfig'

export default {
  name: 'SysParam',
  data() {
    return {
      // 表格数据
      configList: [],
      // 类型数据字典
      typeOptions: [],
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        configName: undefined,
        configKey: undefined,
        configType: undefined,
        // 日期范围
        dateRange: []
      },
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // form 表单：新增或者修改时试用
      form: {},
      // 表单校验
      rules: {
        configName: [
          { required: true, message: '参数名称不能为空', trigger: 'blur' }
        ],
        configKey: [
          { required: true, message: '参数键名不能为空', trigger: 'blur' }
        ],
        configValue: [
          { required: true, message: '参数键值不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      getSysConfigList(this.queryParams).then(response => {
        this.configList = response.data.list
        this.total = response.data.total
        this.loading = false
      }
      )
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageNum = 1
      this.getList()
    },
    // 表单重置
    reset() {
      this.form = {
        configId: undefined,
        configName: undefined,
        configKey: undefined,
        configValue: undefined,
        configType: '1',
        remark: undefined
      }
      this.$refs['queryForm'].resetFields()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.queryParams = {
        pageNum: 1,
        pageSize: 10,
        configName: undefined,
        configKey: undefined,
        configType: undefined,
        // 日期范围
        dateRange: []
      }
      // 指定表单属性值重置
      this.$refs['queryForm'].resetFields()
      this.handleQuery()
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.configId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加参数'
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const configId = row.configId || this.ids
      getConfig(configId).then(response => {
        this.form = {
          configId: response.data.sysConfig.configId,
          configName: response.data.sysConfig.configName,
          configKey: response.data.sysConfig.configKey,
          configValue: response.data.sysConfig.configValue,
          configType: '' + response.data.sysConfig.configType,
          remark: response.data.sysConfig.remark
        }
        this.open = true
        this.title = '修改参数'
      })
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.configId !== undefined) {
            updateConfig(this.form).then(response => {
              if (response.code === 0) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            console.log(this.form)
            addConfig(this.form).then(response => {
              if (response.code === 0) {
                this.msgSuccess('新增成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const configIds = row.configId || this.ids
      this.$confirm('是否确认删除参数编号为"' + configIds + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delConfig(configIds)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() { })
    },
    /** 导出按钮操作 */
    handleExport() {
      const queryParams = this.queryParams
      this.$confirm('是否确认导出所有参数数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return exportConfig(queryParams)
      }).then(response => {
        this.download(response.msg)
      }).catch(function() { })
    }
  }
}
</script>

