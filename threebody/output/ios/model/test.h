
//
//这是一个自动生成的model文件，不要对该文件做任何修改
//作者：shanks
//时间：2015-08-31 23:56:35
//

#import "JSONModel.h"
#import "Home.h"


@interface Test : JSONModel <PMSellerModel>
///这是String类型的数组
@property (nonatomic, strong) NSArray *nick;
@property (nonatomic, strong) NSArray<Home> *home;
@property (nonatomic, assign) int id;
@property (nonatomic, retain) NSString *name;

@end
@protocol Test <NSObject>

@end