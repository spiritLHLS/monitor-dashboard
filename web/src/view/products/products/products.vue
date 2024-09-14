<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
            :disabled-date="time => searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>
        <template v-if="showAllQuery">
          <div class="row">
            <el-form-item label="TAG" prop="tag">
              <el-input v-model="searchInfo.tag" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="CPU" prop="cpu">
              <el-input v-model="searchInfo.cpu" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="内存" prop="memory">
              <el-input v-model="searchInfo.memory" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="磁盘" prop="disk">
              <el-input v-model="searchInfo.disk" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="流量" prop="traffic">
              <el-input v-model="searchInfo.traffic" placeholder="搜索条件" />
            </el-form-item>
          </div>
          <div class="row">
            <el-form-item label="端口" prop="portSpeed">
              <el-input v-model="searchInfo.portSpeed" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="地点" prop="location">
              <el-input v-model="searchInfo.location" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="价格" prop="price">
              <el-input v-model="searchInfo.price" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="其他" prop="additional">
              <el-input v-model="searchInfo.additional" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="链接" prop="url">
              <el-input v-model="searchInfo.url" placeholder="搜索条件" />
            </el-form-item>
          </div>
          <div class="row">
            <el-form-item label="计费类型" prop="billingType">
              <el-input v-model="searchInfo.billingType" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="历史库存" prop="oldStock">
              <el-input v-model.number="searchInfo.oldStock" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="现有库存" prop="stock">
              <el-input v-model.number="searchInfo.stock" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="爬虫间隔" prop="intervals">
              <el-input v-model.number="searchInfo.intervals" placeholder="搜索条件" />
            </el-form-item>
          </div>
          <div class="row">
            <el-form-item label="消息编号" prop="messageId">
              <el-input v-model="searchInfo.messageId" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="推送间隔" prop="pushIntervals">
              <el-input v-model.number="searchInfo.pushIntervals" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="多重检测" prop="multiCheck">
              <el-input v-model="searchInfo.multiCheck" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="推送时间" prop="pushTime">
              <el-date-picker v-model="searchInfo.pushTime" type="datetime" placeholder="搜索条件"></el-date-picker>
            </el-form-item>
          </div>
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery = true"
            v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery = false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="onDelete">删除</el-button>
        <!-- 修改历史库存对话框 -->
        <el-dialog v-model="modifyOldStockDialogVisible" title="修改历史库存" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的历史库存">
              <el-input v-model.number="modifyOldStock" placeholder="请输入新的历史库存"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyOldStockDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmmodifyOldStock">确定</el-button>
          </div>
        </el-dialog>
        <!-- 修改爬取间隔对话框 -->
        <el-dialog v-model="modifyIntervalsDialogVisible" title="修改爬取间隔" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的爬取间隔">
              <el-input v-model.number="modifyIntervals" placeholder="请输入新的爬取间隔"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyIntervalsDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmModifyIntervals">确定</el-button>
          </div>
        </el-dialog>
        <!-- 修改推送间隔对话框 -->
        <el-dialog v-model="modifyPushIntervalsDialogVisible" title="修改推送间隔" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的推送间隔">
              <el-input v-model.number="modifyPushIntervals" placeholder="请输入新的推送间隔"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyPushIntervalsDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmmodifyPushIntervals">确定</el-button>
          </div>
        </el-dialog>
        <!-- 修改计费类型对话框 -->
        <el-dialog v-model="modifyBillingTypeDialogVisible" title="修改计费类型" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的计费类型">
              <el-input v-model.trim="modifyBillingType" placeholder="请输入新的计费类型"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyBillingTypeDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmmodifyBillingType">确定</el-button>
          </div>
        </el-dialog>
        <!-- 修改其他信息对话框 -->
        <el-dialog v-model="modifyAdditionalDialogVisible" title="修改其他信息" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的其他信息">
              <el-input v-model.trim="modifyAdditional" placeholder="请输入新的其他信息"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyAdditionalDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmmodifyAdditional">确定</el-button>
          </div>
        </el-dialog>
        <!-- 修改频道TAG对话框 -->
        <el-dialog v-model="modifyTGTAGDialogVisible" title="修改TGTAG" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的TGTAG">
              <el-input v-model.trim="modifyTGTAG" placeholder="请输入新的TGTAG"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyTGTAGDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmmodifyTGTAG">确定</el-button>
          </div>
        </el-dialog>
        <!-- 修改MessageID对话框 -->
        <el-dialog v-model="modifyMessageIDDialogVisible" title="修改MessageID" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的MessageID">
              <el-input v-model.trim="modifyMessageID" placeholder="请输入新的MessageID"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyMessageIDDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmmodifyMessageID">确定</el-button>
          </div>
        </el-dialog>
        <!-- 修改是否重复检测对话框 -->
        <el-dialog v-model="modifyMultiCheckDialogVisible" title="修改是否重复检测" width="30%">
          <el-form label-width="120px">
            <el-form-item label="是否重复检测">
              <el-input v-model.trim="modifyMultiCheck" placeholder="是则填重复检测次数，否则填0"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyMultiCheckDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmmodifyMultiCheck">确定</el-button>
          </div>
        </el-dialog>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateOldStock">改历史库存</el-button>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateIntervals">改爬取间隔</el-button>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updatePushIntervals">改推送间隔</el-button>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateBillingType">改计费类型</el-button>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateAdditional">改其他信息</el-button>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateTGTAG">改TGTAG</el-button>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateMessageID">改MessageID</el-button>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateMultiCheck">改重复检测次数</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="TAG" prop="tag" width="90" />
        <el-table-column align="left" label="CPU" prop="cpu" width="90" />
        <el-table-column align="left" label="内存" prop="memory" width="90" />
        <el-table-column align="left" label="磁盘" prop="disk" width="90" />
        <el-table-column align="left" label="流量" prop="traffic" width="90" />
        <el-table-column align="left" label="端口" prop="portSpeed" width="90" />
        <el-table-column align="left" label="地点" prop="location" width="90" />
        <el-table-column align="left" label="价格" prop="price" width="90" />
        <el-table-column label="其他" prop="additional" width="200">
          <template #default="scope">
            [富文本内容]
          </template>
        </el-table-column>
        <el-table-column align="left" label="链接" prop="url" width="90" />
        <el-table-column align="left" label="计费类型" prop="billingType" width="90" />
        <el-table-column align="left" label="历史库存" prop="oldStock" width="90" />
        <el-table-column align="left" label="现有库存" prop="stock" width="90" />
        <el-table-column align="left" label="重复检测" prop="multiCheck" width="90" />
        <el-table-column align="left" label="爬虫间隔" prop="intervals" width="90" />
        <el-table-column align="left" label="消息ID" prop="messageId" width="90" />
        <el-table-column align="left" label="推送间隔" prop="pushIntervals" width="90" />
        <el-table-column align="left" label="推送时间" prop="pushTime" width="180">
          <template #default="scope">{{ formatDate(scope.row.pushTime) }}</template>
        </el-table-column>
        <!-- <el-table-column align="left" label="创建者" prop="createdBy" width="90" />
          <el-table-column align="left" label="更新者" prop="updatedBy" width="90" />
          <el-table-column align="left" label="删除者" prop="deletedBy" width="90" /> -->
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon
                style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看详情</el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateProductsFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="TAG:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true" placeholder="请输入TGTAG" />
        </el-form-item>
        <el-form-item label="CPU:" prop="cpu">
          <el-input v-model="formData.cpu" :clearable="true" placeholder="请输入CPU数量" />
        </el-form-item>
        <el-form-item label="内存:" prop="memory">
          <el-input v-model="formData.memory" :clearable="true" placeholder="请输入内存大小" />
        </el-form-item>
        <el-form-item label="磁盘:" prop="disk">
          <el-input v-model="formData.disk" :clearable="true" placeholder="请输入磁盘大小" />
        </el-form-item>
        <el-form-item label="流量:" prop="traffic">
          <el-input v-model="formData.traffic" :clearable="true" placeholder="请输入流量大小" />
        </el-form-item>
        <el-form-item label="端口:" prop="portSpeed">
          <el-input v-model="formData.portSpeed" :clearable="true" placeholder="请输入端口速度信息" />
        </el-form-item>
        <el-form-item label="地点:" prop="location">
          <el-input v-model="formData.location" :clearable="true" placeholder="请输入地点" />
        </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input v-model="formData.price" :clearable="true" placeholder="请输入价格" />
        </el-form-item>
        <el-form-item label="其他:" prop="additional">
          <RichEdit v-model="formData.additional" />
        </el-form-item>
        <el-form-item label="链接:" prop="url">
          <el-input v-model="formData.url" :clearable="true" placeholder="请输入商品链接" />
        </el-form-item>
        <el-form-item label="计费类型:" prop="billingType">
          <el-input v-model="formData.billingType" :clearable="true" placeholder="请输入计费类型" />
        </el-form-item>
        <el-form-item label="历史库存:" prop="oldStock">
          <el-input v-model.number="formData.oldStock" :clearable="true" placeholder="请输入历史库存" />
        </el-form-item>
        <el-form-item label="现有库存:" prop="stock">
          <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入现有库存" />
        </el-form-item>
        <el-form-item label="重复检测:" prop="multiCheck">
          <el-input v-model.number="formData.multiCheck" :clearable="true" placeholder="请输入重复检测" />
        </el-form-item>
        <el-form-item label="爬虫间隔:" prop="intervals">
          <el-input v-model.number="formData.intervals" :clearable="true" placeholder="请输入爬虫间隔" />
        </el-form-item>
        <el-form-item label="消息编号:" prop="messageId">
          <el-input v-model="formData.messageId" :clearable="true" placeholder="请输入消息编号" />
        </el-form-item>
        <el-form-item label="推送间隔:" prop="pushIntervals">
          <el-input v-model.number="formData.pushIntervals" :clearable="true" placeholder="请输入推送间隔，间隔为-1时不进行推送" />
        </el-form-item>
        <el-form-item label="推送时间:" prop="pushTime">
          <el-date-picker v-model="formData.pushTime" type="date" style="width:100%" placeholder="选择日期"
            :clearable="true" />
        </el-form-item>
        <!-- <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入创建者" />
        </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入更新者" />
        </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入删除者" />
        </el-form-item> -->
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="TAG">
          {{ detailFrom.tag }}
        </el-descriptions-item>
        <el-descriptions-item label="CPU">
          {{ detailFrom.cpu }}
        </el-descriptions-item>
        <el-descriptions-item label="内存">
          {{ detailFrom.memory }}
        </el-descriptions-item>
        <el-descriptions-item label="磁盘">
          {{ detailFrom.disk }}
        </el-descriptions-item>
        <el-descriptions-item label="流量">
          {{ detailFrom.traffic }}
        </el-descriptions-item>
        <el-descriptions-item label="端口">
          {{ detailFrom.portSpeed }}
        </el-descriptions-item>
        <el-descriptions-item label="地点">
          {{ detailFrom.location }}
        </el-descriptions-item>
        <el-descriptions-item label="价格">
          {{ detailFrom.price }}
        </el-descriptions-item>
        <el-descriptions-item label="其他">
          {{ detailFrom.additional }}
        </el-descriptions-item>
        <el-descriptions-item label="链接">
          {{ detailFrom.url }}
        </el-descriptions-item>
        <el-descriptions-item label="计费类型">
          {{ detailFrom.billingType }}
        </el-descriptions-item>
        <el-descriptions-item label="历史库存">
          {{ detailFrom.oldStock }}
        </el-descriptions-item>
        <el-descriptions-item label="现有库存">
          {{ detailFrom.stock }}
        </el-descriptions-item>
        <el-descriptions-item label="重复检测">
          {{ detailFrom.multiCheck }}
        </el-descriptions-item>
        <el-descriptions-item label="爬虫间隔">
          {{ detailFrom.intervals }}
        </el-descriptions-item>
        <el-descriptions-item label="消息编号">
          {{ detailFrom.messageId }}
        </el-descriptions-item>
        <el-descriptions-item label="推送间隔">
          {{ detailFrom.pushIntervals }}
        </el-descriptions-item>
        <el-descriptions-item label="推送时间">
          {{ detailFrom.pushTime }}
        </el-descriptions-item>
        <!-- <el-descriptions-item label="创建者">
          {{ detailFrom.createdBy }}
        </el-descriptions-item>
        <el-descriptions-item label="更新者">
          {{ detailFrom.updatedBy }}
        </el-descriptions-item>
        <el-descriptions-item label="删除者">
          {{ detailFrom.deletedBy }}
        </el-descriptions-item> -->
      </el-descriptions>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  createProducts,
  deleteProducts,
  deleteProductsByIds,
  updateProducts,
  findProducts,
  getProductsList
} from '@/api/products/products'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'Products'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  tag: '',
  cpu: '',
  memory: '',
  disk: '',
  traffic: '',
  portSpeed: '',
  location: '',
  price: '',
  additional: '',
  url: '',
  billingType: '',
  oldStock: 0,
  stock: 0,
  multiCheck: 0,
  intervals: 0,
  messageId: '',
  pushIntervals: 0,
  pushTime: new Date(),
  createdBy: undefined,
  updatedBy: undefined,
  deletedBy: undefined,
})



// 验证规则
const rule = reactive({
  tag: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  url: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  stock: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  intervals: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  pushIntervals: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
  pushTime: [{
    validator: (rule, value, callback) => {
      if (searchInfo.value.startPushTime && !searchInfo.value.endPushTime) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startPushTime && searchInfo.value.endPushTime) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startPushTime && searchInfo.value.endPushTime && (searchInfo.value.startPushTime.getTime() === searchInfo.value.endPushTime.getTime() || searchInfo.value.startPushTime.getTime() > searchInfo.value.endPushTime.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change'
  }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    // console.log('Search Info in onSubmit:', JSON.stringify(searchInfo.value)) // 添加日志
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getProductsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    // console.log('Updated tableData:', JSON.stringify(tableData.value)) // 添加日志
  } else {
    console.error('API Error:', table.message) // 添加错误日志
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteProductsFunc(row)
  })
}

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
    const res = await deleteProductsByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateProductsFunc = async (row) => {
  const res = await findProducts({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteProductsFunc = async (row) => {
  const res = await deleteProducts({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    tag: '',
    cpu: '',
    memory: '',
    disk: '',
    traffic: '',
    portSpeed: '',
    location: '',
    price: '',
    additional: '',
    url: '',
    billingType: '',
    oldStock: 0,
    stock: 0,
    multiCheck: 0,
    intervals: 0,
    messageId: '',
    pushIntervals: 0,
    pushTime: new Date(),
    createdBy: '',
    updatedBy: '',
    deletedBy: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createProducts(formData.value)
        break
      case 'update':
        res = await updateProducts(formData.value)
        break
      default:
        res = await createProducts(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

// 批量更新推送间隔、其他信息
const modifyIntervalsDialogVisible = ref(false)
const modifyPushIntervalsDialogVisible = ref(false)
const modifyAdditionalDialogVisible = ref(false)
const modifyBillingTypeDialogVisible = ref(false)
const modifyTGTAGDialogVisible = ref(false)
const modifyMessageIDDialogVisible = ref(false)
const modifyOldStockDialogVisible = ref(false)
const modifyMultiCheckDialogVisible = ref(false)

// 更新所选记录的爬取间隔的方法
const updateIntervals = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyIntervalsDialogVisible.value = true
}

// 更新所选记录的推送间隔的方法
const updatePushIntervals = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyPushIntervalsDialogVisible.value = true
}

// 更新所选记录的计费类型的方法
const updateBillingType = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyBillingTypeDialogVisible.value = true
}

// 更新所选记录的其他信息的方法
const updateAdditional = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyAdditionalDialogVisible.value = true
}

// 更新所选记录的TGTAG的方法
const updateTGTAG = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyTGTAGDialogVisible.value = true
}

// 更新所选记录的MessageID的方法
const updateMessageID = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyMessageIDDialogVisible.value = true
}

// 更新所选记录的MultiCheck的方法
const updateMultiCheck = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyMultiCheckDialogVisible.value = true
}

// 更新所选记录的OldStock的方法
const updateOldStock = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyOldStockDialogVisible.value = true
}

// 修改的新值
const modifyIntervals = ref(0)
const modifyPushIntervals = ref(0)
const modifyBillingType = ref("")
const modifyAdditional = ref("")
const modifyTGTAG = ref("")
const modifyMessageID = ref("")
const modifyOldStock = ref(0)
const modifyMultiCheck = ref(0)

// 确认修改爬取间隔
const confirmModifyIntervals = async () => {
  // 获取用户输入的新值
  const newIntervals = modifyIntervals.value
  // 调用更新函数
  await updateSelectedRecords(newIntervals, 'intervals')
  // 关闭对话框
  modifyIntervalsDialogVisible.value = false
}

// 确认修改推送间隔
const confirmmodifyPushIntervals = async () => {
  // 获取用户输入的新值
  const newIntervals = modifyPushIntervals.value
  // 调用更新函数
  await updateSelectedRecords(newIntervals, 'pushIntervals')
  // 关闭对话框
  modifyPushIntervalsDialogVisible.value = false
}

// 确认修改计费类型
const confirmmodifyBillingType = async () => {
  // 获取用户输入的新值
  const newbillingType = modifyBillingType.value
  // 调用更新函数
  await updateSelectedRecords(newbillingType, 'billingType')
  // 关闭对话框
  modifyBillingTypeDialogVisible.value = false
}

// 确认修改其他信息
const confirmmodifyAdditional = async () => {
  // 获取用户输入的新值
  const newadditional = modifyAdditional.value
  // 调用更新函数
  await updateSelectedRecords(newadditional, 'additional')
  // 关闭对话框
  modifyAdditionalDialogVisible.value = false
}

// 确认修改TGTAG
const confirmmodifyTGTAG = async () => {
  // 获取用户输入的新值
  const newtgtag = modifyTGTAG.value
  // 调用更新函数
  await updateSelectedRecords(newtgtag, 'tag')
  // 关闭对话框
  modifyTGTAGDialogVisible.value = false
}

// 确认修改MessageID
const confirmmodifyMessageID = async () => {
  // 获取用户输入的新值
  const newMessageID = modifyMessageID.value
  // 调用更新函数
  await updateSelectedRecords(newMessageID, 'messageId')
  // 关闭对话框
  modifyMessageIDDialogVisible.value = false
}

// 确认修改MultiCheck
const confirmmodifyMultiCheck = async () => {
  // 获取用户输入的新值
  const newMultiCheckString = modifyMultiCheck.value
  // 将字符串转换为整数
  const newMultiCheck = parseInt(newMultiCheckString, 10)
  // 检查转换是否成功（不是 NaN）
  if (isNaN(newMultiCheck)) {
    // 如果转换失败，可以显示错误消息
    console.error('Invalid input: not a number')
    // 可以添加用户提示，例如：
    // alert('请输入有效的数字')
    return // 终止函数执行
  }
  // 调用更新函数，传入转换后的整数值
  await updateSelectedRecords(newMultiCheck, 'multiCheck')
  // 关闭对话框
  modifyMultiCheckDialogVisible.value = false
}


// 确认修改OldStock
const confirmmodifyOldStock = async () => {
  // 获取用户输入的新值
  const newOldStock = modifyOldStock.value
  // 调用更新函数
  await updateSelectedRecords(newOldStock, 'oldStock')
  // 关闭对话框
  modifyOldStockDialogVisible.value = false
}

// 更新选中记录的指定字段值
const updateSelectedRecords = async (newValue, field) => {
  // 获取选中记录的ID列表
  const IDs = multipleSelection.value.map(record => record.ID)
  // 构造更新数据
  const updatedRecords = multipleSelection.value.map(record => ({
    ...record,
    [field]: newValue
  }))
  // 调用更新接口，这里调用你的更新接口函数
  const res = await Promise.all(updatedRecords.map(record => updateProducts(record)))
  // 处理响应
  handleUpdateResponse(res)
}

// 处理来自API更新的响应
const handleUpdateResponse = (responses) => {
  const isSuccess = responses.every(res => res.code === 0)
  if (isSuccess) {
    ElMessage({
      type: 'success',
      message: '修改成功'
    })
    getTableData() // 刷新表格数据
  } else {
    ElMessage({
      type: 'error',
      message: '修改失败，请稍后重试'
    })
  }
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findProducts({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    tag: '',
    cpu: '',
    memory: '',
    disk: '',
    traffic: '',
    portSpeed: '',
    location: '',
    price: '',
    additional: '',
    url: '',
    billingType: '',
    oldStock: 0,
    stock: 0,
    intervals: 0,
    messageId: '',
    pushIntervals: 0,
    pushTime: new Date(),
  }
}


</script>

<style>
.row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1px;
  /* 行与行之间的间距 */
}

.row .el-form-item {
  flex: 1;
  /* 使每个表单项占据相等的空间 */
  margin-right: 10px;
  /* 表单项之间的间距 */
}

.el-table .cell {
  /* 文本超出容器部分隐藏 */
  overflow: hidden;
  /* 超出部分使用省略号代替 */
  text-overflow: ellipsis;
  /* 不换行 */
  white-space: nowrap;
}

.demo-form-inline .el-input {
  width: 160px;
  /* 输入框的宽度 */
}
</style>