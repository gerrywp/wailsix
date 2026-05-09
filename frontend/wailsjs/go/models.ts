export namespace main {
	
	export class Config {
	    // Go type: struct { BaseURL string "json:\"baseURL\""; Timeout int "json:\"timeout\"" }
	    esbAPI: any;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.esbAPI = this.convertValues(source["esbAPI"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace models {
	
	export class IPqcCheckTask {
	    ObjectRrn: number;
	    OrgRrn: number;
	    IsActive: string;
	    // Go type: time
	    Created: any;
	    CreatedBy: string;
	    // Go type: time
	    Updated: any;
	    UpdatedBy: string;
	    LockVersion: number;
	    CheckCode: string;
	    LotId: string;
	    SampleIds: string;
	    CheckType: number;
	    CheckReason: string;
	    CheckTimes: number;
	    ProductNumberVer: string;
	    TriggerStep: string;
	    CheckName: string;
	    TakePosition: string;
	    WeekCode: number;
	    Layers: number;
	    LayerLevel: string;
	    CustomerCode: string;
	    State: string;
	    LotNumber: string;
	    CheckNumber: number;
	    AccNumber: number;
	    Unit: string;
	    RejNumber: number;
	    DecisionResult: string;
	    OriDecisionResult: string;
	    DecisionRemark: string;
	    OriCheckCode: string;
	    CustomerModel: string;
	    // Go type: time
	    PostTime: any;
	    // Go type: time
	    AuditTime: any;
	    Auditby: string;
	    Remark: string;
	    Stage: string;
	    NextAuditName: string;
	    // Go type: time
	    NextAuditTime: any;
	    IsRecheck: string;
	    DefectType: string;
	    RecheckRemark: string;
	    Attachment: string;
	    DefectLevel: string;
	    Source: number;
	    EquipmentId: string;
	    CheckNameRrn: number;
	    SonLotId: string;
	    Ipqctype: string;
	    Checkwhy: number;
	    Entermark: string;
	    DeleteRemark: string;
	    NcmCode: string;
	    NcnCode: string;
	    Entryby: string;
	
	    static createFrom(source: any = {}) {
	        return new IPqcCheckTask(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ObjectRrn = source["ObjectRrn"];
	        this.OrgRrn = source["OrgRrn"];
	        this.IsActive = source["IsActive"];
	        this.Created = this.convertValues(source["Created"], null);
	        this.CreatedBy = source["CreatedBy"];
	        this.Updated = this.convertValues(source["Updated"], null);
	        this.UpdatedBy = source["UpdatedBy"];
	        this.LockVersion = source["LockVersion"];
	        this.CheckCode = source["CheckCode"];
	        this.LotId = source["LotId"];
	        this.SampleIds = source["SampleIds"];
	        this.CheckType = source["CheckType"];
	        this.CheckReason = source["CheckReason"];
	        this.CheckTimes = source["CheckTimes"];
	        this.ProductNumberVer = source["ProductNumberVer"];
	        this.TriggerStep = source["TriggerStep"];
	        this.CheckName = source["CheckName"];
	        this.TakePosition = source["TakePosition"];
	        this.WeekCode = source["WeekCode"];
	        this.Layers = source["Layers"];
	        this.LayerLevel = source["LayerLevel"];
	        this.CustomerCode = source["CustomerCode"];
	        this.State = source["State"];
	        this.LotNumber = source["LotNumber"];
	        this.CheckNumber = source["CheckNumber"];
	        this.AccNumber = source["AccNumber"];
	        this.Unit = source["Unit"];
	        this.RejNumber = source["RejNumber"];
	        this.DecisionResult = source["DecisionResult"];
	        this.OriDecisionResult = source["OriDecisionResult"];
	        this.DecisionRemark = source["DecisionRemark"];
	        this.OriCheckCode = source["OriCheckCode"];
	        this.CustomerModel = source["CustomerModel"];
	        this.PostTime = this.convertValues(source["PostTime"], null);
	        this.AuditTime = this.convertValues(source["AuditTime"], null);
	        this.Auditby = source["Auditby"];
	        this.Remark = source["Remark"];
	        this.Stage = source["Stage"];
	        this.NextAuditName = source["NextAuditName"];
	        this.NextAuditTime = this.convertValues(source["NextAuditTime"], null);
	        this.IsRecheck = source["IsRecheck"];
	        this.DefectType = source["DefectType"];
	        this.RecheckRemark = source["RecheckRemark"];
	        this.Attachment = source["Attachment"];
	        this.DefectLevel = source["DefectLevel"];
	        this.Source = source["Source"];
	        this.EquipmentId = source["EquipmentId"];
	        this.CheckNameRrn = source["CheckNameRrn"];
	        this.SonLotId = source["SonLotId"];
	        this.Ipqctype = source["Ipqctype"];
	        this.Checkwhy = source["Checkwhy"];
	        this.Entermark = source["Entermark"];
	        this.DeleteRemark = source["DeleteRemark"];
	        this.NcmCode = source["NcmCode"];
	        this.NcnCode = source["NcnCode"];
	        this.Entryby = source["Entryby"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

