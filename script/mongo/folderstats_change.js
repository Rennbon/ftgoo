//数据库结构替换
db.folderStatistics.find({}).noCursorTimeout().forEach((ft)=>{
    ft.ds.forEach((o)=>{
        db.folderStatisticsNew.insert({
            fid:ft.fid,
            ct:ft.ct,
            
            dt:o.dt,//1
            at:o.at,//2
            u:o.u,//3
            un:o.un,//4
            ua:o.ua,//5
            uu:o.uu,//6
            t:o.t,//7
            tu:o.tu,//8
            tc:o.tc,//9
            c:o.c,//10
            cn:o.cn,//11
            ca:o.ca,//12
            cu:o.cu,//13
            ma:o.ma,//14
            cha:o.cha,//15
            nts:o.nts,//16
            cts:o.cts,//17
            tbd:o.tbd//18
        })
        
    })
    
})