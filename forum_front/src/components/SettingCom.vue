<template>
    <div>
        <el-card>
            <div>

                <h3>头像</h3>
                <el-image :src="userInfo.avatar" class="myAvatar">
                    <div slot="placeholder" class="image-slot">
                        加载中<span class="dot">...</span>
                    </div>
                </el-image>
                <el-upload
                        ref="upload"
                        class="upload-demo"
                        :action="avatarAction"
                        :headers="myHeader"
                        accept=".jpg,.jpeg,.JPG,.JPEG,.PNG,.png"
                        :on-success="uploadSuccess"
                        :before-upload="beforeUplaod"
                        :show-file-list="false"
                        :on-exceed="handleExceed"
                        list-type="picture"
                >
                    <el-button size="small" type="success">修改头像</el-button>
                    <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过1M</div>
                </el-upload>
            </div>
        </el-card>
        <el-card >
            用户名：
            <span v-text="userInfo.username"></span><br>
        </el-card>
        <el-card >
            性别：
            <span v-text="userInfo.gender == 0 ? '男' : '女'"></span>
        </el-card>
        <el-card >
            注册时间：
            <span v-text="userInfo.createTime"></span>
        </el-card>
        <el-card>
            <h3>修改密码</h3>
            <el-form :model="passwordForm" :rules="passwordRules" ref="passwordForm" label-width="100px">
                <el-form-item label="旧密码" prop="oldPassword">
                    <el-input v-model="passwordForm.oldPassword" type="password" placeholder="请输入旧密码"></el-input>
                </el-form-item>
                <el-form-item label="新密码" prop="newPassword">
                    <el-input v-model="passwordForm.newPassword" type="password" placeholder="请输入新密码"></el-input>
                </el-form-item>
                <el-form-item label="确认密码" prop="confirmPassword">
                    <el-input v-model="passwordForm.confirmPassword" type="password" placeholder="请再次输入新密码"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="changePassword">提交</el-button>
                    <el-button @click="resetPasswordForm">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>

<script>
    export default {
        name: "SettingCom",
        data(){
            return{
                //头像网络资源地址
                uploadPath:this.$axios.defaults.baseURL,
                // picturePath:'http://47.115.88.155',
                userInfo:{},
                isMine:false,
                likeCount:'',

                //头像上传地址
                avatarAction: this.$axios.defaults.baseURL.replace(/\/$/, '')+'/user/avatar',
                myHeader:{'Authorization':sessionStorage.getItem("JWT_TOKEN")},

                //修改密码表单
                passwordForm: {
                    oldPassword: '',
                    newPassword: '',
                    confirmPassword: ''
                },
                //密码验证规则
                passwordRules: {
                    oldPassword: [
                        { required: true, message: '请输入旧密码', trigger: 'blur' }
                    ],
                    newPassword: [
                        { required: true, message: '请输入新密码', trigger: 'blur' },
                        { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
                    ],
                    confirmPassword: [
                        { required: true, message: '请再次输入新密码', trigger: 'blur' },
                        { validator: this.validateConfirmPassword, trigger: 'blur' }
                    ]
                }
            }
        },
        created() {
            const _this = this
            //请求页面资源
            this.$axios({
                method:'get',
                url:'/user/getInfo'
            }).then(function(res){
                if (res.data.code == 200){
                    const userData = res.data.data;
                    // 拼接完整的头像URL
                    if (userData.avatar && !userData.avatar.startsWith('http')) {
                        userData.avatar = _this.$axios.defaults.baseURL.replace(/\/$/, '') + userData.avatar;
                    }
                    _this.userInfo = userData;
                    // 同时更新 store 中的用户信息
                    _this.$store.commit("setUserInfo", userData);
                }else{
                    _this.fail(res.data.msg);
                    _this.router.push('/login')
                }

            }).catch(function(error){
                console.log(error);
            });
        },
        methods:{
            fail(msg) {
                this.$message.error(msg);
            },
            handleExceed(file,fileList){
                this.$message.error("只能上传一张图片")
            },
            uploadSuccess(response, file, fileList){
                console.log(response);
                if (response.code === 200) {
                    this.userInfo.avatar = this.$axios.defaults.baseURL.replace(/\/$/, '') + response.data.avatar;
                    //更新用户个人信息
                    this.$store.commit("setUserInfo", this.userInfo);
                    //强制页面重新渲染
                    this.$forceUpdate();
                } else {
                    this.$message.error(response.message || "上传失败");
                }
            },
            beforeUplaod(file){
                const isLt1M = file.size / 1024 / 1024 < 1 ;
                if (!isLt1M) {
                    this.$message.error('上传头像图片大小不能超过 1MB!');
                }
                return isLt1M;
            },
            //验证确认密码
            validateConfirmPassword(rule, value, callback) {
                if (value === '') {
                    callback(new Error('请再次输入新密码'));
                } else if (value !== this.passwordForm.newPassword) {
                    callback(new Error('两次输入的密码不一致'));
                } else {
                    callback();
                }
            },
            //修改密码
            changePassword() {
                const _this = this;
                this.$refs.passwordForm.validate((valid) => {
                    if (valid) {
                        this.$axios({
                            method: 'post',
                            url: '/user/resetPass',
                            data: {
                                oldPassword: this.passwordForm.oldPassword,
                                newPassword: this.passwordForm.newPassword
                            }
                        }).then(function(res){
                            if (res.data.code === 200) {
                                _this.$message.success('修改密码成功');
                                _this.resetPasswordForm();
                            } else {
                                _this.$message.error(res.data.message || '修改密码失败');
                            }
                        }).catch(function(error){
                            console.log(error);
                            _this.$message.error('修改密码失败');
                        });
                    } else {
                        return false;
                    }
                });
            },
            //重置密码表单
            resetPasswordForm() {
                this.$refs.passwordForm.resetFields();
            }
        }
    }
</script>

<style scoped>
    .myAvatar{
        width: 200px;
        height: 200px;
    }

     .avatar-uploader .el-upload {
         border: 1px dashed #d9d9d9;
         border-radius: 6px;
         cursor: pointer;
         position: relative;
         overflow: hidden;
     }
    .avatar-uploader .el-upload:hover {
        border-color: #409EFF;
    }
    .avatar-uploader-icon {
        font-size: 28px;
        color: #8c939d;
        width: 178px;
        height: 178px;
        line-height: 178px;
        text-align: center;
    }
    .avatar {
        width: 178px;
        height: 178px;
        display: block;
    }

</style>