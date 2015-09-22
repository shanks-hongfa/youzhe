
//
//这是一个自动生成的model文件，不要对该文件做任何修改
//作者：shanks
//时间：2015-08-31 23:56:35
//

#import "JSONModel.h"


@interface Home : JSONModel <PMSellerModel>
@property (nonatomic, assign) int id;
@property (nonatomic, retain) NSString *name;
///这是String类型的数组
@property (nonatomic, strong) NSArray *nick;

@end
@protocol Home <NSObject>

@end